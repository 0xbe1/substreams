#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Manifest {
    #[prost(string, tag="1")]
    pub spec_version: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub description: ::prost::alloc::string::String,
    #[prost(message, repeated, tag="3")]
    pub modules: ::prost::alloc::vec::Vec<Module>,
    #[prost(bytes="vec", repeated, tag="4")]
    pub modules_code: ::prost::alloc::vec::Vec<::prost::alloc::vec::Vec<u8>>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Module {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
    #[prost(message, repeated, tag="6")]
    pub inputs: ::prost::alloc::vec::Vec<module::Input>,
    #[prost(message, optional, tag="7")]
    pub output: ::core::option::Option<module::Output>,
    #[prost(uint64, tag="8")]
    pub start_block: u64,
    #[prost(oneof="module::Kind", tags="2, 3")]
    pub kind: ::core::option::Option<module::Kind>,
    #[prost(oneof="module::Code", tags="4, 5")]
    pub code: ::core::option::Option<module::Code>,
}
/// Nested message and enum types in `Module`.
pub mod module {
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct WasmCode {
        #[prost(string, tag="4")]
        pub r#type: ::prost::alloc::string::String,
        #[prost(uint32, tag="5")]
        pub index: u32,
        #[prost(string, tag="6")]
        pub entrypoint: ::prost::alloc::string::String,
    }
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct NativeCode {
        #[prost(string, tag="5")]
        pub entrypoint: ::prost::alloc::string::String,
    }
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct KindMap {
        #[prost(string, tag="1")]
        pub output_type: ::prost::alloc::string::String,
    }
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct KindStore {
        /// The `update_policy` determines the functions available to mutate the store
        /// (like `set()`, `set_if_not_exists()` or `sum()`, etc..) in
        /// order to ensure that parallel operations are possible and deterministic
        ///
        /// Say a store cumulates keys from block 0 to 1M, and a second store
        /// cumulates keys from block 1M to 2M. When we want to use this
        /// store as a dependency for a downstream module, we will merge the
        /// two stores according to this policy.
        #[prost(enumeration="kind_store::UpdatePolicy", tag="1")]
        pub update_policy: i32,
        #[prost(string, tag="2")]
        pub value_type: ::prost::alloc::string::String,
    }
    /// Nested message and enum types in `KindStore`.
    pub mod kind_store {
        #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
        #[repr(i32)]
        pub enum UpdatePolicy {
            Unset = 0,
            /// Provides a store where you can `set()` keys, and the latest key wins
            Replace = 1,
            /// Provides a store where you can `set_if_not_exists()` keys, and the first key wins
            Ignore = 2,
            /// Provides a store where you can `sum_*()` keys, where two stores merge by summing its values.
            Sum = 3,
            /// Provides a store where you can `min_*()` keys, where two stores merge by leaving the minimum value.
            Min = 4,
            /// Provides a store where you can `max_*()` keys, where two stores merge by leaving the maximum value.
            Max = 5,
        }
    }
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct Input {
        #[prost(oneof="input::Input", tags="1, 2, 3")]
        pub input: ::core::option::Option<input::Input>,
    }
    /// Nested message and enum types in `Input`.
    pub mod input {
        #[derive(Clone, PartialEq, ::prost::Message)]
        pub struct Source {
            /// ex: "sf.ethereum.type.v1.Block"
            #[prost(string, tag="1")]
            pub r#type: ::prost::alloc::string::String,
        }
        #[derive(Clone, PartialEq, ::prost::Message)]
        pub struct Map {
            /// ex: "block_to_pairs"
            #[prost(string, tag="1")]
            pub module_name: ::prost::alloc::string::String,
        }
        #[derive(Clone, PartialEq, ::prost::Message)]
        pub struct Store {
            #[prost(string, tag="1")]
            pub module_name: ::prost::alloc::string::String,
            #[prost(enumeration="store::Mode", tag="2")]
            pub mode: i32,
        }
        /// Nested message and enum types in `Store`.
        pub mod store {
            #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
            #[repr(i32)]
            pub enum Mode {
                Unset = 0,
                Get = 1,
                Deltas = 2,
            }
        }
        #[derive(Clone, PartialEq, ::prost::Oneof)]
        pub enum Input {
            #[prost(message, tag="1")]
            Source(Source),
            #[prost(message, tag="2")]
            Map(Map),
            #[prost(message, tag="3")]
            Store(Store),
        }
    }
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct Output {
        #[prost(string, tag="1")]
        pub r#type: ::prost::alloc::string::String,
    }
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Kind {
        #[prost(message, tag="2")]
        KindMap(KindMap),
        #[prost(message, tag="3")]
        KindStore(KindStore),
    }
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Code {
        #[prost(message, tag="4")]
        WasmCode(WasmCode),
        #[prost(message, tag="5")]
        NativeCode(NativeCode),
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Clock {
    #[prost(string, tag="1")]
    pub id: ::prost::alloc::string::String,
    #[prost(uint64, tag="2")]
    pub number: u64,
    #[prost(message, optional, tag="3")]
    pub timestamp: ::core::option::Option<::prost_types::Timestamp>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Request {
    #[prost(int64, tag="1")]
    pub start_block_num: i64,
    #[prost(string, tag="2")]
    pub start_cursor: ::prost::alloc::string::String,
    #[prost(uint64, tag="3")]
    pub stop_block_num: u64,
    #[prost(enumeration="ForkStep", repeated, tag="4")]
    pub fork_steps: ::prost::alloc::vec::Vec<i32>,
    #[prost(string, tag="5")]
    pub irreversibility_condition: ::prost::alloc::string::String,
    #[prost(message, optional, tag="6")]
    pub manifest: ::core::option::Option<Manifest>,
    #[prost(string, repeated, tag="7")]
    pub output_modules: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(string, repeated, tag="8")]
    pub initial_store_snapshot_for_modules: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Response {
    #[prost(oneof="response::Message", tags="1, 2, 3, 4")]
    pub message: ::core::option::Option<response::Message>,
}
/// Nested message and enum types in `Response`.
pub mod response {
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Message {
        /// Progress of data preparation, before sending in the stream of `data` events.
        #[prost(message, tag="1")]
        Progress(super::ModulesProgress),
        #[prost(message, tag="2")]
        SnapshotData(super::InitialSnapshotData),
        #[prost(message, tag="3")]
        SnapshotComplete(super::InitialSnapshotComplete),
        #[prost(message, tag="4")]
        Data(super::BlockScopedData),
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct InitialSnapshotComplete {
    #[prost(string, tag="1")]
    pub cursor: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct InitialSnapshotData {
    #[prost(string, tag="1")]
    pub module_name: ::prost::alloc::string::String,
    #[prost(message, optional, tag="2")]
    pub deltas: ::core::option::Option<StoreDeltas>,
    #[prost(uint64, tag="4")]
    pub sent_keys: u64,
    #[prost(uint64, tag="3")]
    pub total_keys: u64,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BlockScopedData {
    #[prost(message, repeated, tag="1")]
    pub outputs: ::prost::alloc::vec::Vec<ModuleOutput>,
    #[prost(message, optional, tag="3")]
    pub clock: ::core::option::Option<Clock>,
    #[prost(enumeration="ForkStep", tag="6")]
    pub step: i32,
    #[prost(string, tag="10")]
    pub cursor: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ModuleOutput {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
    #[prost(string, repeated, tag="4")]
    pub logs: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(oneof="module_output::Data", tags="2, 3")]
    pub data: ::core::option::Option<module_output::Data>,
}
/// Nested message and enum types in `ModuleOutput`.
pub mod module_output {
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Data {
        #[prost(message, tag="2")]
        MapOutput(::prost_types::Any),
        #[prost(message, tag="3")]
        StoreDeltas(super::StoreDeltas),
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ModulesProgress {
    #[prost(message, repeated, tag="1")]
    pub modules: ::prost::alloc::vec::Vec<ModuleProgress>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ModuleProgress {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
    #[prost(message, repeated, tag="2")]
    pub processed_ranges: ::prost::alloc::vec::Vec<BlockRange>,
    #[prost(uint64, tag="3")]
    pub total_bytes_read: u64,
    #[prost(uint64, tag="4")]
    pub total_bytes_written: u64,
    #[prost(bool, tag="7")]
    pub failed: bool,
    #[prost(string, tag="8")]
    pub failure_reason: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BlockRange {
    #[prost(uint64, tag="1")]
    pub start_block: u64,
    #[prost(uint64, tag="2")]
    pub end_block: u64,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct StoreDeltas {
    #[prost(message, repeated, tag="1")]
    pub deltas: ::prost::alloc::vec::Vec<StoreDelta>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct StoreDelta {
    #[prost(enumeration="store_delta::Operation", tag="1")]
    pub operation: i32,
    #[prost(uint64, tag="2")]
    pub ordinal: u64,
    #[prost(string, tag="3")]
    pub key: ::prost::alloc::string::String,
    #[prost(bytes="vec", tag="4")]
    pub old_value: ::prost::alloc::vec::Vec<u8>,
    #[prost(bytes="vec", tag="5")]
    pub new_value: ::prost::alloc::vec::Vec<u8>,
}
/// Nested message and enum types in `StoreDelta`.
pub mod store_delta {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum Operation {
        Unset = 0,
        Create = 1,
        Update = 2,
        Delete = 3,
    }
}
//table.1
// create f.1 o.10 n.99
// update f.1 o.99 n.200
// update f.2 o.abc n.xyz
// update f.1 o.200 n.400

//table.1
// create f.1 o.10 n.400 f.2 o.abc n.xyz

//table.1
// update f.1 o.10 n.99
// update f.1 o.99 n.200
// update f.2 o.abc n.xyz
// update f.1 o.200 n.400

//table.1
// update f.1 o.10 n.400 f.2 o.abc n.xyz

#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Output {
    #[prost(uint64, tag="1")]
    pub block_num: u64,
    #[prost(string, tag="2")]
    pub block_id: ::prost::alloc::string::String,
    #[prost(message, optional, tag="4")]
    pub timestamp: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(message, optional, tag="10")]
    pub value: ::core::option::Option<::prost_types::Any>,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum ForkStep {
    StepUnknown = 0,
    /// Block is new head block of the chain, that is linear with the previous block
    StepNew = 1,
    /// Block is now forked and should be undone, it's not the head block of the chain anymore
    StepUndo = 2,
    /// Block is now irreversible and can be committed to (finality is chain specific, see chain documentation for more details)
    StepIrreversible = 4,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct DatabaseChanges {
    #[prost(message, repeated, tag="1")]
    pub table_changes: ::prost::alloc::vec::Vec<TableChange>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct TableChange {
    #[prost(string, tag="1")]
    pub table: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub pk: ::prost::alloc::string::String,
    #[prost(uint64, tag="3")]
    pub block_num: u64,
    #[prost(uint64, tag="4")]
    pub ordinal: u64,
    #[prost(enumeration="table_change::Operation", tag="5")]
    pub operation: i32,
    #[prost(message, repeated, tag="6")]
    pub fields: ::prost::alloc::vec::Vec<Field>,
}
/// Nested message and enum types in `TableChange`.
pub mod table_change {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum Operation {
        Unset = 0,
        Create = 1,
        Update = 2,
        Delete = 3,
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Field {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub new_value: ::prost::alloc::string::String,
    #[prost(string, tag="3")]
    pub old_value: ::prost::alloc::string::String,
}
