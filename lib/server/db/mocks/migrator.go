// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	sync "sync"

	db "github.com/hyperledger/fabric-ca/lib/server/db"
)

type Migrator struct {
	CommitStub        func() error
	commitMutex       sync.RWMutex
	commitArgsForCall []struct {
	}
	commitReturns struct {
		result1 error
	}
	commitReturnsOnCall map[int]struct {
		result1 error
	}
	MigrateAffiliationsTableStub        func() error
	migrateAffiliationsTableMutex       sync.RWMutex
	migrateAffiliationsTableArgsForCall []struct {
	}
	migrateAffiliationsTableReturns struct {
		result1 error
	}
	migrateAffiliationsTableReturnsOnCall map[int]struct {
		result1 error
	}
	MigrateCertificatesTableStub        func() error
	migrateCertificatesTableMutex       sync.RWMutex
	migrateCertificatesTableArgsForCall []struct {
	}
	migrateCertificatesTableReturns struct {
		result1 error
	}
	migrateCertificatesTableReturnsOnCall map[int]struct {
		result1 error
	}
	MigrateCredentialsTableStub        func() error
	migrateCredentialsTableMutex       sync.RWMutex
	migrateCredentialsTableArgsForCall []struct {
	}
	migrateCredentialsTableReturns struct {
		result1 error
	}
	migrateCredentialsTableReturnsOnCall map[int]struct {
		result1 error
	}
	MigrateNoncesTableStub        func() error
	migrateNoncesTableMutex       sync.RWMutex
	migrateNoncesTableArgsForCall []struct {
	}
	migrateNoncesTableReturns struct {
		result1 error
	}
	migrateNoncesTableReturnsOnCall map[int]struct {
		result1 error
	}
	MigrateRAInfoTableStub        func() error
	migrateRAInfoTableMutex       sync.RWMutex
	migrateRAInfoTableArgsForCall []struct {
	}
	migrateRAInfoTableReturns struct {
		result1 error
	}
	migrateRAInfoTableReturnsOnCall map[int]struct {
		result1 error
	}
	MigrateUsersTableStub        func() error
	migrateUsersTableMutex       sync.RWMutex
	migrateUsersTableArgsForCall []struct {
	}
	migrateUsersTableReturns struct {
		result1 error
	}
	migrateUsersTableReturnsOnCall map[int]struct {
		result1 error
	}
	RollbackStub        func() error
	rollbackMutex       sync.RWMutex
	rollbackArgsForCall []struct {
	}
	rollbackReturns struct {
		result1 error
	}
	rollbackReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Migrator) Commit() error {
	fake.commitMutex.Lock()
	ret, specificReturn := fake.commitReturnsOnCall[len(fake.commitArgsForCall)]
	fake.commitArgsForCall = append(fake.commitArgsForCall, struct {
	}{})
	fake.recordInvocation("Commit", []interface{}{})
	fake.commitMutex.Unlock()
	if fake.CommitStub != nil {
		return fake.CommitStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.commitReturns
	return fakeReturns.result1
}

func (fake *Migrator) CommitCallCount() int {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return len(fake.commitArgsForCall)
}

func (fake *Migrator) CommitCalls(stub func() error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = stub
}

func (fake *Migrator) CommitReturns(result1 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	fake.commitReturns = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) CommitReturnsOnCall(i int, result1 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	if fake.commitReturnsOnCall == nil {
		fake.commitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.commitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateAffiliationsTable() error {
	fake.migrateAffiliationsTableMutex.Lock()
	ret, specificReturn := fake.migrateAffiliationsTableReturnsOnCall[len(fake.migrateAffiliationsTableArgsForCall)]
	fake.migrateAffiliationsTableArgsForCall = append(fake.migrateAffiliationsTableArgsForCall, struct {
	}{})
	fake.recordInvocation("MigrateAffiliationsTable", []interface{}{})
	fake.migrateAffiliationsTableMutex.Unlock()
	if fake.MigrateAffiliationsTableStub != nil {
		return fake.MigrateAffiliationsTableStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.migrateAffiliationsTableReturns
	return fakeReturns.result1
}

func (fake *Migrator) MigrateAffiliationsTableCallCount() int {
	fake.migrateAffiliationsTableMutex.RLock()
	defer fake.migrateAffiliationsTableMutex.RUnlock()
	return len(fake.migrateAffiliationsTableArgsForCall)
}

func (fake *Migrator) MigrateAffiliationsTableCalls(stub func() error) {
	fake.migrateAffiliationsTableMutex.Lock()
	defer fake.migrateAffiliationsTableMutex.Unlock()
	fake.MigrateAffiliationsTableStub = stub
}

func (fake *Migrator) MigrateAffiliationsTableReturns(result1 error) {
	fake.migrateAffiliationsTableMutex.Lock()
	defer fake.migrateAffiliationsTableMutex.Unlock()
	fake.MigrateAffiliationsTableStub = nil
	fake.migrateAffiliationsTableReturns = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateAffiliationsTableReturnsOnCall(i int, result1 error) {
	fake.migrateAffiliationsTableMutex.Lock()
	defer fake.migrateAffiliationsTableMutex.Unlock()
	fake.MigrateAffiliationsTableStub = nil
	if fake.migrateAffiliationsTableReturnsOnCall == nil {
		fake.migrateAffiliationsTableReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.migrateAffiliationsTableReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateCertificatesTable() error {
	fake.migrateCertificatesTableMutex.Lock()
	ret, specificReturn := fake.migrateCertificatesTableReturnsOnCall[len(fake.migrateCertificatesTableArgsForCall)]
	fake.migrateCertificatesTableArgsForCall = append(fake.migrateCertificatesTableArgsForCall, struct {
	}{})
	fake.recordInvocation("MigrateCertificatesTable", []interface{}{})
	fake.migrateCertificatesTableMutex.Unlock()
	if fake.MigrateCertificatesTableStub != nil {
		return fake.MigrateCertificatesTableStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.migrateCertificatesTableReturns
	return fakeReturns.result1
}

func (fake *Migrator) MigrateCertificatesTableCallCount() int {
	fake.migrateCertificatesTableMutex.RLock()
	defer fake.migrateCertificatesTableMutex.RUnlock()
	return len(fake.migrateCertificatesTableArgsForCall)
}

func (fake *Migrator) MigrateCertificatesTableCalls(stub func() error) {
	fake.migrateCertificatesTableMutex.Lock()
	defer fake.migrateCertificatesTableMutex.Unlock()
	fake.MigrateCertificatesTableStub = stub
}

func (fake *Migrator) MigrateCertificatesTableReturns(result1 error) {
	fake.migrateCertificatesTableMutex.Lock()
	defer fake.migrateCertificatesTableMutex.Unlock()
	fake.MigrateCertificatesTableStub = nil
	fake.migrateCertificatesTableReturns = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateCertificatesTableReturnsOnCall(i int, result1 error) {
	fake.migrateCertificatesTableMutex.Lock()
	defer fake.migrateCertificatesTableMutex.Unlock()
	fake.MigrateCertificatesTableStub = nil
	if fake.migrateCertificatesTableReturnsOnCall == nil {
		fake.migrateCertificatesTableReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.migrateCertificatesTableReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateCredentialsTable() error {
	fake.migrateCredentialsTableMutex.Lock()
	ret, specificReturn := fake.migrateCredentialsTableReturnsOnCall[len(fake.migrateCredentialsTableArgsForCall)]
	fake.migrateCredentialsTableArgsForCall = append(fake.migrateCredentialsTableArgsForCall, struct {
	}{})
	fake.recordInvocation("MigrateCredentialsTable", []interface{}{})
	fake.migrateCredentialsTableMutex.Unlock()
	if fake.MigrateCredentialsTableStub != nil {
		return fake.MigrateCredentialsTableStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.migrateCredentialsTableReturns
	return fakeReturns.result1
}

func (fake *Migrator) MigrateCredentialsTableCallCount() int {
	fake.migrateCredentialsTableMutex.RLock()
	defer fake.migrateCredentialsTableMutex.RUnlock()
	return len(fake.migrateCredentialsTableArgsForCall)
}

func (fake *Migrator) MigrateCredentialsTableCalls(stub func() error) {
	fake.migrateCredentialsTableMutex.Lock()
	defer fake.migrateCredentialsTableMutex.Unlock()
	fake.MigrateCredentialsTableStub = stub
}

func (fake *Migrator) MigrateCredentialsTableReturns(result1 error) {
	fake.migrateCredentialsTableMutex.Lock()
	defer fake.migrateCredentialsTableMutex.Unlock()
	fake.MigrateCredentialsTableStub = nil
	fake.migrateCredentialsTableReturns = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateCredentialsTableReturnsOnCall(i int, result1 error) {
	fake.migrateCredentialsTableMutex.Lock()
	defer fake.migrateCredentialsTableMutex.Unlock()
	fake.MigrateCredentialsTableStub = nil
	if fake.migrateCredentialsTableReturnsOnCall == nil {
		fake.migrateCredentialsTableReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.migrateCredentialsTableReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateNoncesTable() error {
	fake.migrateNoncesTableMutex.Lock()
	ret, specificReturn := fake.migrateNoncesTableReturnsOnCall[len(fake.migrateNoncesTableArgsForCall)]
	fake.migrateNoncesTableArgsForCall = append(fake.migrateNoncesTableArgsForCall, struct {
	}{})
	fake.recordInvocation("MigrateNoncesTable", []interface{}{})
	fake.migrateNoncesTableMutex.Unlock()
	if fake.MigrateNoncesTableStub != nil {
		return fake.MigrateNoncesTableStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.migrateNoncesTableReturns
	return fakeReturns.result1
}

func (fake *Migrator) MigrateNoncesTableCallCount() int {
	fake.migrateNoncesTableMutex.RLock()
	defer fake.migrateNoncesTableMutex.RUnlock()
	return len(fake.migrateNoncesTableArgsForCall)
}

func (fake *Migrator) MigrateNoncesTableCalls(stub func() error) {
	fake.migrateNoncesTableMutex.Lock()
	defer fake.migrateNoncesTableMutex.Unlock()
	fake.MigrateNoncesTableStub = stub
}

func (fake *Migrator) MigrateNoncesTableReturns(result1 error) {
	fake.migrateNoncesTableMutex.Lock()
	defer fake.migrateNoncesTableMutex.Unlock()
	fake.MigrateNoncesTableStub = nil
	fake.migrateNoncesTableReturns = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateNoncesTableReturnsOnCall(i int, result1 error) {
	fake.migrateNoncesTableMutex.Lock()
	defer fake.migrateNoncesTableMutex.Unlock()
	fake.MigrateNoncesTableStub = nil
	if fake.migrateNoncesTableReturnsOnCall == nil {
		fake.migrateNoncesTableReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.migrateNoncesTableReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateRAInfoTable() error {
	fake.migrateRAInfoTableMutex.Lock()
	ret, specificReturn := fake.migrateRAInfoTableReturnsOnCall[len(fake.migrateRAInfoTableArgsForCall)]
	fake.migrateRAInfoTableArgsForCall = append(fake.migrateRAInfoTableArgsForCall, struct {
	}{})
	fake.recordInvocation("MigrateRAInfoTable", []interface{}{})
	fake.migrateRAInfoTableMutex.Unlock()
	if fake.MigrateRAInfoTableStub != nil {
		return fake.MigrateRAInfoTableStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.migrateRAInfoTableReturns
	return fakeReturns.result1
}

func (fake *Migrator) MigrateRAInfoTableCallCount() int {
	fake.migrateRAInfoTableMutex.RLock()
	defer fake.migrateRAInfoTableMutex.RUnlock()
	return len(fake.migrateRAInfoTableArgsForCall)
}

func (fake *Migrator) MigrateRAInfoTableCalls(stub func() error) {
	fake.migrateRAInfoTableMutex.Lock()
	defer fake.migrateRAInfoTableMutex.Unlock()
	fake.MigrateRAInfoTableStub = stub
}

func (fake *Migrator) MigrateRAInfoTableReturns(result1 error) {
	fake.migrateRAInfoTableMutex.Lock()
	defer fake.migrateRAInfoTableMutex.Unlock()
	fake.MigrateRAInfoTableStub = nil
	fake.migrateRAInfoTableReturns = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateRAInfoTableReturnsOnCall(i int, result1 error) {
	fake.migrateRAInfoTableMutex.Lock()
	defer fake.migrateRAInfoTableMutex.Unlock()
	fake.MigrateRAInfoTableStub = nil
	if fake.migrateRAInfoTableReturnsOnCall == nil {
		fake.migrateRAInfoTableReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.migrateRAInfoTableReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateUsersTable() error {
	fake.migrateUsersTableMutex.Lock()
	ret, specificReturn := fake.migrateUsersTableReturnsOnCall[len(fake.migrateUsersTableArgsForCall)]
	fake.migrateUsersTableArgsForCall = append(fake.migrateUsersTableArgsForCall, struct {
	}{})
	fake.recordInvocation("MigrateUsersTable", []interface{}{})
	fake.migrateUsersTableMutex.Unlock()
	if fake.MigrateUsersTableStub != nil {
		return fake.MigrateUsersTableStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.migrateUsersTableReturns
	return fakeReturns.result1
}

func (fake *Migrator) MigrateUsersTableCallCount() int {
	fake.migrateUsersTableMutex.RLock()
	defer fake.migrateUsersTableMutex.RUnlock()
	return len(fake.migrateUsersTableArgsForCall)
}

func (fake *Migrator) MigrateUsersTableCalls(stub func() error) {
	fake.migrateUsersTableMutex.Lock()
	defer fake.migrateUsersTableMutex.Unlock()
	fake.MigrateUsersTableStub = stub
}

func (fake *Migrator) MigrateUsersTableReturns(result1 error) {
	fake.migrateUsersTableMutex.Lock()
	defer fake.migrateUsersTableMutex.Unlock()
	fake.MigrateUsersTableStub = nil
	fake.migrateUsersTableReturns = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) MigrateUsersTableReturnsOnCall(i int, result1 error) {
	fake.migrateUsersTableMutex.Lock()
	defer fake.migrateUsersTableMutex.Unlock()
	fake.MigrateUsersTableStub = nil
	if fake.migrateUsersTableReturnsOnCall == nil {
		fake.migrateUsersTableReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.migrateUsersTableReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) Rollback() error {
	fake.rollbackMutex.Lock()
	ret, specificReturn := fake.rollbackReturnsOnCall[len(fake.rollbackArgsForCall)]
	fake.rollbackArgsForCall = append(fake.rollbackArgsForCall, struct {
	}{})
	fake.recordInvocation("Rollback", []interface{}{})
	fake.rollbackMutex.Unlock()
	if fake.RollbackStub != nil {
		return fake.RollbackStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.rollbackReturns
	return fakeReturns.result1
}

func (fake *Migrator) RollbackCallCount() int {
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	return len(fake.rollbackArgsForCall)
}

func (fake *Migrator) RollbackCalls(stub func() error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = stub
}

func (fake *Migrator) RollbackReturns(result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	fake.rollbackReturns = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) RollbackReturnsOnCall(i int, result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	if fake.rollbackReturnsOnCall == nil {
		fake.rollbackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rollbackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Migrator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	fake.migrateAffiliationsTableMutex.RLock()
	defer fake.migrateAffiliationsTableMutex.RUnlock()
	fake.migrateCertificatesTableMutex.RLock()
	defer fake.migrateCertificatesTableMutex.RUnlock()
	fake.migrateCredentialsTableMutex.RLock()
	defer fake.migrateCredentialsTableMutex.RUnlock()
	fake.migrateNoncesTableMutex.RLock()
	defer fake.migrateNoncesTableMutex.RUnlock()
	fake.migrateRAInfoTableMutex.RLock()
	defer fake.migrateRAInfoTableMutex.RUnlock()
	fake.migrateUsersTableMutex.RLock()
	defer fake.migrateUsersTableMutex.RUnlock()
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Migrator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.Migrator = new(Migrator)
