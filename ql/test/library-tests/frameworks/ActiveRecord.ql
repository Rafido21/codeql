import codeql_ruby.controlflow.CfgNodes
import codeql_ruby.frameworks.ActiveRecord

query predicate activeRecordModelClasses(ActiveRecordModelClass cls) { any() }

query predicate activeRecordSqlExecutionRanges(ActiveRecordSqlExecutionRange range) { any() }

query predicate activeRecordModelClassMethodCalls(ActiveRecordModelClassMethodCall call) { any() }

query predicate sqlExecutingMethodCall(SqlExecutingMethodCall call) { any() }
