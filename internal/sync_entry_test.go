package internal

import (
	"sync"
	"testing"

	"github.com/ghosind/go-assert"
)

func TestNewSyncEntry(t *testing.T) {
	a := assert.New(t)

	// Test with int type
	expunged := new(int)
	entry := NewSyncEntry(42, expunged)
	a.NotNilNow(entry)
	a.EqualNow(expunged, entry.expunged)

	// Verify initial value is stored correctly
	val, ok := entry.Load(0)
	a.TrueNow(ok)
	a.EqualNow(42, val)

	// Test with string type
	expungedStr := new(string)
	entryStr := NewSyncEntry("hello", expungedStr)
	a.NotNilNow(entryStr)
	a.EqualNow(expungedStr, entryStr.expunged)

	valStr, ok := entryStr.Load("")
	a.TrueNow(ok)
	a.EqualNow("hello", valStr)
}

func TestSyncEntryLoad(t *testing.T) {
	a := assert.New(t)
	expunged := new(int)

	// Test loading valid value
	entry := NewSyncEntry(42, expunged)
	val, ok := entry.Load(0)
	a.TrueNow(ok)
	a.EqualNow(42, val)

	// Test loading when entry is nil
	entry.p.Store(nil)
	val, ok = entry.Load(99)
	a.NotTrueNow(ok)
	a.EqualNow(99, val) // Should return default value

	// Test loading when entry is expunged
	entry.p.Store(expunged)
	val, ok = entry.Load(99)
	a.NotTrueNow(ok)
	a.EqualNow(99, val) // Should return default value
}

func TestSyncEntryTrySwap(t *testing.T) {
	a := assert.New(t)
	expunged := new(int)

	// Test successful swap
	entry := NewSyncEntry(42, expunged)
	newVal := 100
	oldVal, ok := entry.TrySwap(&newVal)
	a.TrueNow(ok)
	a.EqualNow(42, *oldVal)

	// Verify new value is stored
	val, loaded := entry.Load(0)
	a.TrueNow(loaded)
	a.EqualNow(100, val)

	// Test swap when entry is expunged (should fail)
	entry.p.Store(expunged)
	anotherVal := 200
	oldVal, ok = entry.TrySwap(&anotherVal)
	a.NotTrueNow(ok)
	a.NilNow(oldVal)
}

func TestSyncEntryTrySwapConcurrent(t *testing.T) {
	a := assert.New(t)
	expunged := new(int)
	entry := NewSyncEntry(42, expunged)

	const numGoroutines = 10
	var wg sync.WaitGroup
	results := make([]bool, numGoroutines)
	values := make([]*int, numGoroutines)
	oldValues := make([]*int, numGoroutines)

	// Create unique values for each goroutine
	for i := 0; i < numGoroutines; i++ {
		val := i + 100
		values[i] = &val
	}

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(idx int) {
			defer wg.Done()
			old, ok := entry.TrySwap(values[idx])
			results[idx] = ok
			oldValues[idx] = old
		}(i)
	}

	wg.Wait()

	// All swaps should succeed in this case
	successCount := 0
	for _, result := range results {
		if result {
			successCount++
		}
	}
	a.EqualNow(numGoroutines, successCount)

	// Verify that the final value is one of the values we set
	finalVal, ok := entry.Load(0)
	a.TrueNow(ok)
	found := false
	for i := 0; i < numGoroutines; i++ {
		if finalVal == i+100 {
			found = true
			break
		}
	}
	a.TrueNow(found)
}

func TestSyncEntryDelete(t *testing.T) {
	a := assert.New(t)
	expunged := new(int)

	// Test successful delete
	entry := NewSyncEntry(42, expunged)
	deletedVal, ok := entry.Delete()
	a.TrueNow(ok)
	a.EqualNow(42, *deletedVal)

	// Verify entry is now nil
	val, loaded := entry.Load(99)
	a.NotTrueNow(loaded)
	a.EqualNow(99, val)

	// Test delete when already nil (should fail)
	deletedVal, ok = entry.Delete()
	a.NotTrueNow(ok)
	a.NilNow(deletedVal)

	// Test delete when expunged (should fail)
	entry.p.Store(expunged)
	deletedVal, ok = entry.Delete()
	a.NotTrueNow(ok)
	a.NilNow(deletedVal)
}

func TestSyncEntryDeleteConcurrent(t *testing.T) {
	a := assert.New(t)
	expunged := new(int)
	entry := NewSyncEntry(42, expunged)

	const numGoroutines = 10
	var wg sync.WaitGroup
	results := make([]bool, numGoroutines)
	deletedValues := make([]*int, numGoroutines)

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(idx int) {
			defer wg.Done()
			val, ok := entry.Delete()
			results[idx] = ok
			deletedValues[idx] = val
		}(i)
	}

	wg.Wait()

	// Exactly one delete should succeed
	successCount := 0
	for i, result := range results {
		if result {
			successCount++
			a.NotNilNow(deletedValues[i])
			a.EqualNow(42, *deletedValues[i])
		} else {
			a.NilNow(deletedValues[i])
		}
	}
	a.EqualNow(1, successCount)
}

func TestSyncEntryUnexpungeLocked(t *testing.T) {
	a := assert.New(t)
	expunged := new(int)

	// Test unexpunge when expunged
	entry := NewSyncEntry(42, expunged)
	entry.p.Store(expunged) // Set to expunged state
	result := entry.UnexpungeLocked()
	a.TrueNow(result)

	// Verify it's now nil (unexpunged)
	p := entry.p.Load()
	a.NilNow(p)

	// Test unexpunge when not expunged (should fail)
	val := 100
	entry.p.Store(&val)
	result = entry.UnexpungeLocked()
	a.NotTrueNow(result)

	// Verify value is unchanged
	p = entry.p.Load()
	a.NotNilNow(p)
	a.EqualNow(100, *p)
}

func TestSyncEntrySwapLocked(t *testing.T) {
	a := assert.New(t)
	expunged := new(int)

	// Test swap with existing value
	entry := NewSyncEntry(42, expunged)
	newVal := 100
	oldVal := entry.SwapLocked(&newVal)
	a.NotNilNow(oldVal)
	a.EqualNow(42, *oldVal)

	// Verify new value is stored
	p := entry.p.Load()
	a.NotNilNow(p)
	a.EqualNow(100, *p)

	// Test swap with nil
	oldVal = entry.SwapLocked(nil)
	a.NotNilNow(oldVal)
	a.EqualNow(100, *oldVal)

	// Verify it's now nil
	p = entry.p.Load()
	a.NilNow(p)
}

func TestSyncEntryTryExpungeLocked(t *testing.T) {
	a := assert.New(t)
	expunged := new(int)

	// Test expunge when nil
	entry := NewSyncEntry(42, expunged)
	entry.p.Store(nil) // Set to nil first
	result := entry.TryExpungeLocked()
	a.TrueNow(result)

	// Verify it's now expunged
	p := entry.p.Load()
	a.EqualNow(expunged, p)

	// Test expunge when already expunged (should return true)
	result = entry.TryExpungeLocked()
	a.TrueNow(result)

	// Test expunge when has value (should fail)
	val := 100
	entry.p.Store(&val)
	result = entry.TryExpungeLocked()
	a.NotTrueNow(result)

	// Verify value is unchanged
	p = entry.p.Load()
	a.NotNilNow(p)
	a.EqualNow(100, *p)
}

func TestSyncEntryTryExpungeLockedConcurrent(t *testing.T) {
	a := assert.New(t)
	expunged := new(int)
	entry := NewSyncEntry(42, expunged)
	entry.p.Store(nil) // Start with nil

	const numGoroutines = 10
	var wg sync.WaitGroup
	results := make([]bool, numGoroutines)

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(idx int) {
			defer wg.Done()
			results[idx] = entry.TryExpungeLocked()
		}(i)
	}

	wg.Wait()

	// All should succeed since they all try to expunge nil
	for _, result := range results {
		a.TrueNow(result)
	}

	// Verify final state is expunged
	p := entry.p.Load()
	a.EqualNow(expunged, p)
}

func TestSyncEntryComplexWorkflow(t *testing.T) {
	a := assert.New(t)
	expunged := new(string)

	// Create entry with initial value
	entry := NewSyncEntry("initial", expunged)

	// Load initial value
	val, ok := entry.Load("")
	a.TrueNow(ok)
	a.EqualNow("initial", val)

	// Swap to new value
	newVal := "swapped"
	oldVal, ok := entry.TrySwap(&newVal)
	a.TrueNow(ok)
	a.EqualNow("initial", *oldVal)

	// Verify new value
	val, ok = entry.Load("")
	a.TrueNow(ok)
	a.EqualNow("swapped", val)

	// Delete the entry
	deletedVal, ok := entry.Delete()
	a.TrueNow(ok)
	a.EqualNow("swapped", *deletedVal)

	// Try to load (should fail)
	val, ok = entry.Load("default")
	a.NotTrueNow(ok)
	a.EqualNow("default", val)

	// Try expunge (should succeed since it's nil)
	success := entry.TryExpungeLocked()
	a.TrueNow(success)

	// Try to swap expunged entry (should fail)
	anotherVal := "failed"
	oldVal, ok = entry.TrySwap(&anotherVal)
	a.NotTrueNow(ok)
	a.NilNow(oldVal)

	// Unexpunge
	success = entry.UnexpungeLocked()
	a.TrueNow(success)

	// Now we can swap again
	finalVal := "final"
	entry.p.Store(&finalVal)
	val, ok = entry.Load("")
	a.TrueNow(ok)
	a.EqualNow("final", val)
}

func TestSyncReadOnlyStruct(t *testing.T) {
	a := assert.New(t)

	// Test SyncReadOnly struct creation and field access
	expunged := new(int)
	entry1 := NewSyncEntry(42, expunged)
	entry2 := NewSyncEntry(100, expunged)

	syncReadOnly := SyncReadOnly[string, int]{
		M: map[string]*SyncEntry[int]{
			"key1": entry1,
			"key2": entry2,
		},
		Amended: true,
	}

	a.NotNilNow(syncReadOnly.M)
	a.EqualNow(2, len(syncReadOnly.M))
	a.TrueNow(syncReadOnly.Amended)

	// Test accessing entries through the map
	entry, exists := syncReadOnly.M["key1"]
	a.TrueNow(exists)
	a.NotNilNow(entry)

	val, ok := entry.Load(0)
	a.TrueNow(ok)
	a.EqualNow(42, val)
}
