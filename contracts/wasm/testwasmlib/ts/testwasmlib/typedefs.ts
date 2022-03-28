// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ArrayOfImmutableAddress extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getAddress(index: u32): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.index(index));
	}
}

export class ImmutableAddressArray extends ArrayOfImmutableAddress {
}

export class ArrayOfMutableAddress extends wasmtypes.ScProxy {

	appendAddress(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getAddress(index: u32): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.index(index));
	}
}

export class MutableAddressArray extends ArrayOfMutableAddress {
}

export class MapAddressToImmutableAddress extends wasmtypes.ScProxy {

	getAddress(key: wasmtypes.ScAddress): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.key(wasmtypes.addressToBytes(key)));
	}
}

export class ImmutableAddressMap extends MapAddressToImmutableAddress {
}

export class MapAddressToMutableAddress extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getAddress(key: wasmtypes.ScAddress): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.key(wasmtypes.addressToBytes(key)));
	}
}

export class MutableAddressMap extends MapAddressToMutableAddress {
}

export class ArrayOfImmutableAgentID extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getAgentID(index: u32): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.index(index));
	}
}

export class ImmutableAgentIDArray extends ArrayOfImmutableAgentID {
}

export class ArrayOfMutableAgentID extends wasmtypes.ScProxy {

	appendAgentID(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getAgentID(index: u32): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.index(index));
	}
}

export class MutableAgentIDArray extends ArrayOfMutableAgentID {
}

export class MapAgentIDToImmutableAgentID extends wasmtypes.ScProxy {

	getAgentID(key: wasmtypes.ScAgentID): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.key(wasmtypes.agentIDToBytes(key)));
	}
}

export class ImmutableAgentIDMap extends MapAgentIDToImmutableAgentID {
}

export class MapAgentIDToMutableAgentID extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getAgentID(key: wasmtypes.ScAgentID): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.key(wasmtypes.agentIDToBytes(key)));
	}
}

export class MutableAgentIDMap extends MapAgentIDToMutableAgentID {
}

export class ArrayOfImmutableBytes extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getBytes(index: u32): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.index(index));
	}
}

export class ImmutableBytesArray extends ArrayOfImmutableBytes {
}

export class ArrayOfMutableBytes extends wasmtypes.ScProxy {

	appendBytes(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getBytes(index: u32): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.index(index));
	}
}

export class MutableBytesArray extends ArrayOfMutableBytes {
}

export class MapBytesToImmutableBytes extends wasmtypes.ScProxy {

	getBytes(key: u8[]): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.key(wasmtypes.bytesToBytes(key)));
	}
}

export class ImmutableBytesMap extends MapBytesToImmutableBytes {
}

export class MapBytesToMutableBytes extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getBytes(key: u8[]): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.key(wasmtypes.bytesToBytes(key)));
	}
}

export class MutableBytesMap extends MapBytesToMutableBytes {
}

export class ArrayOfImmutableChainID extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getChainID(index: u32): wasmtypes.ScImmutableChainID {
		return new wasmtypes.ScImmutableChainID(this.proxy.index(index));
	}
}

export class ImmutableChainIDArray extends ArrayOfImmutableChainID {
}

export class ArrayOfMutableChainID extends wasmtypes.ScProxy {

	appendChainID(): wasmtypes.ScMutableChainID {
		return new wasmtypes.ScMutableChainID(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getChainID(index: u32): wasmtypes.ScMutableChainID {
		return new wasmtypes.ScMutableChainID(this.proxy.index(index));
	}
}

export class MutableChainIDArray extends ArrayOfMutableChainID {
}

export class MapChainIDToImmutableChainID extends wasmtypes.ScProxy {

	getChainID(key: wasmtypes.ScChainID): wasmtypes.ScImmutableChainID {
		return new wasmtypes.ScImmutableChainID(this.proxy.key(wasmtypes.chainIDToBytes(key)));
	}
}

export class ImmutableChainIDMap extends MapChainIDToImmutableChainID {
}

export class MapChainIDToMutableChainID extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getChainID(key: wasmtypes.ScChainID): wasmtypes.ScMutableChainID {
		return new wasmtypes.ScMutableChainID(this.proxy.key(wasmtypes.chainIDToBytes(key)));
	}
}

export class MutableChainIDMap extends MapChainIDToMutableChainID {
}

export class ArrayOfImmutableColor extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getColor(index: u32): wasmtypes.ScImmutableColor {
		return new wasmtypes.ScImmutableColor(this.proxy.index(index));
	}
}

export class ImmutableColorArray extends ArrayOfImmutableColor {
}

export class ArrayOfMutableColor extends wasmtypes.ScProxy {

	appendColor(): wasmtypes.ScMutableColor {
		return new wasmtypes.ScMutableColor(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getColor(index: u32): wasmtypes.ScMutableColor {
		return new wasmtypes.ScMutableColor(this.proxy.index(index));
	}
}

export class MutableColorArray extends ArrayOfMutableColor {
}

export class MapColorToImmutableColor extends wasmtypes.ScProxy {

	getColor(key: wasmtypes.ScTokenID): wasmtypes.ScImmutableColor {
		return new wasmtypes.ScImmutableColor(this.proxy.key(wasmtypes.colorToBytes(key)));
	}
}

export class ImmutableColorMap extends MapColorToImmutableColor {
}

export class MapColorToMutableColor extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getColor(key: wasmtypes.ScTokenID): wasmtypes.ScMutableColor {
		return new wasmtypes.ScMutableColor(this.proxy.key(wasmtypes.colorToBytes(key)));
	}
}

export class MutableColorMap extends MapColorToMutableColor {
}

export class ArrayOfImmutableHash extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getHash(index: u32): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.index(index));
	}
}

export class ImmutableHashArray extends ArrayOfImmutableHash {
}

export class ArrayOfMutableHash extends wasmtypes.ScProxy {

	appendHash(): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getHash(index: u32): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.index(index));
	}
}

export class MutableHashArray extends ArrayOfMutableHash {
}

export class MapHashToImmutableHash extends wasmtypes.ScProxy {

	getHash(key: wasmtypes.ScHash): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.key(wasmtypes.hashToBytes(key)));
	}
}

export class ImmutableHashMap extends MapHashToImmutableHash {
}

export class MapHashToMutableHash extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getHash(key: wasmtypes.ScHash): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.key(wasmtypes.hashToBytes(key)));
	}
}

export class MutableHashMap extends MapHashToMutableHash {
}

export class ArrayOfImmutableHname extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getHname(index: u32): wasmtypes.ScImmutableHname {
		return new wasmtypes.ScImmutableHname(this.proxy.index(index));
	}
}

export class ImmutableHnameArray extends ArrayOfImmutableHname {
}

export class ArrayOfMutableHname extends wasmtypes.ScProxy {

	appendHname(): wasmtypes.ScMutableHname {
		return new wasmtypes.ScMutableHname(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getHname(index: u32): wasmtypes.ScMutableHname {
		return new wasmtypes.ScMutableHname(this.proxy.index(index));
	}
}

export class MutableHnameArray extends ArrayOfMutableHname {
}

export class MapHnameToImmutableHname extends wasmtypes.ScProxy {

	getHname(key: wasmtypes.ScHname): wasmtypes.ScImmutableHname {
		return new wasmtypes.ScImmutableHname(this.proxy.key(wasmtypes.hnameToBytes(key)));
	}
}

export class ImmutableHnameMap extends MapHnameToImmutableHname {
}

export class MapHnameToMutableHname extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getHname(key: wasmtypes.ScHname): wasmtypes.ScMutableHname {
		return new wasmtypes.ScMutableHname(this.proxy.key(wasmtypes.hnameToBytes(key)));
	}
}

export class MutableHnameMap extends MapHnameToMutableHname {
}

export class ArrayOfImmutableInt32 extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getInt32(index: u32): wasmtypes.ScImmutableInt32 {
		return new wasmtypes.ScImmutableInt32(this.proxy.index(index));
	}
}

export class ImmutableInt32Array extends ArrayOfImmutableInt32 {
}

export class ArrayOfMutableInt32 extends wasmtypes.ScProxy {

	appendInt32(): wasmtypes.ScMutableInt32 {
		return new wasmtypes.ScMutableInt32(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getInt32(index: u32): wasmtypes.ScMutableInt32 {
		return new wasmtypes.ScMutableInt32(this.proxy.index(index));
	}
}

export class MutableInt32Array extends ArrayOfMutableInt32 {
}

export class MapInt32ToImmutableInt32 extends wasmtypes.ScProxy {

	getInt32(key: i32): wasmtypes.ScImmutableInt32 {
		return new wasmtypes.ScImmutableInt32(this.proxy.key(wasmtypes.int32ToBytes(key)));
	}
}

export class ImmutableInt32Map extends MapInt32ToImmutableInt32 {
}

export class MapInt32ToMutableInt32 extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getInt32(key: i32): wasmtypes.ScMutableInt32 {
		return new wasmtypes.ScMutableInt32(this.proxy.key(wasmtypes.int32ToBytes(key)));
	}
}

export class MutableInt32Map extends MapInt32ToMutableInt32 {
}

export class ArrayOfImmutableInt64 extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getInt64(index: u32): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.index(index));
	}
}

export class ImmutableInt64Array extends ArrayOfImmutableInt64 {
}

export class ArrayOfMutableInt64 extends wasmtypes.ScProxy {

	appendInt64(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getInt64(index: u32): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.index(index));
	}
}

export class MutableInt64Array extends ArrayOfMutableInt64 {
}

export class MapInt64ToImmutableInt64 extends wasmtypes.ScProxy {

	getInt64(key: i64): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.key(wasmtypes.int64ToBytes(key)));
	}
}

export class ImmutableInt64Map extends MapInt64ToImmutableInt64 {
}

export class MapInt64ToMutableInt64 extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getInt64(key: i64): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.key(wasmtypes.int64ToBytes(key)));
	}
}

export class MutableInt64Map extends MapInt64ToMutableInt64 {
}

export class MapInt32ToImmutableLocation extends wasmtypes.ScProxy {

	getLocation(key: i32): sc.ImmutableLocation {
		return new sc.ImmutableLocation(this.proxy.key(wasmtypes.int32ToBytes(key)));
	}
}

export class ImmutableLongitude extends MapInt32ToImmutableLocation {
}

export class MapInt32ToMutableLocation extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getLocation(key: i32): sc.MutableLocation {
		return new sc.MutableLocation(this.proxy.key(wasmtypes.int32ToBytes(key)));
	}
}

export class MutableLongitude extends MapInt32ToMutableLocation {
}

export class ArrayOfImmutableRequestID extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getRequestID(index: u32): wasmtypes.ScImmutableRequestID {
		return new wasmtypes.ScImmutableRequestID(this.proxy.index(index));
	}
}

export class ImmutableRequestIDArray extends ArrayOfImmutableRequestID {
}

export class ArrayOfMutableRequestID extends wasmtypes.ScProxy {

	appendRequestID(): wasmtypes.ScMutableRequestID {
		return new wasmtypes.ScMutableRequestID(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getRequestID(index: u32): wasmtypes.ScMutableRequestID {
		return new wasmtypes.ScMutableRequestID(this.proxy.index(index));
	}
}

export class MutableRequestIDArray extends ArrayOfMutableRequestID {
}

export class MapRequestIDToImmutableRequestID extends wasmtypes.ScProxy {

	getRequestID(key: wasmtypes.ScRequestID): wasmtypes.ScImmutableRequestID {
		return new wasmtypes.ScImmutableRequestID(this.proxy.key(wasmtypes.requestIDToBytes(key)));
	}
}

export class ImmutableRequestIDMap extends MapRequestIDToImmutableRequestID {
}

export class MapRequestIDToMutableRequestID extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getRequestID(key: wasmtypes.ScRequestID): wasmtypes.ScMutableRequestID {
		return new wasmtypes.ScMutableRequestID(this.proxy.key(wasmtypes.requestIDToBytes(key)));
	}
}

export class MutableRequestIDMap extends MapRequestIDToMutableRequestID {
}

export class ArrayOfImmutableString extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getString(index: u32): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.index(index));
	}
}

export class ImmutableStringArray extends ArrayOfImmutableString {
}

export class ArrayOfMutableString extends wasmtypes.ScProxy {

	appendString(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getString(index: u32): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.index(index));
	}
}

export class MutableStringArray extends ArrayOfMutableString {
}

export class MapStringToImmutableString extends wasmtypes.ScProxy {

	getString(key: string): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class ImmutableStringMap extends MapStringToImmutableString {
}

export class MapStringToMutableString extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getString(key: string): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class MutableStringMap extends MapStringToMutableString {
}

export class ArrayOfImmutableUint64 extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getUint64(index: u32): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.index(index));
	}
}

export class ImmutableUint64Array extends ArrayOfImmutableUint64 {
}

export class ArrayOfMutableUint64 extends wasmtypes.ScProxy {

	appendUint64(): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getUint64(index: u32): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.index(index));
	}
}

export class MutableUint64Array extends ArrayOfMutableUint64 {
}

export class MapUint64ToImmutableUint64 extends wasmtypes.ScProxy {

	getUint64(key: u64): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.key(wasmtypes.uint64ToBytes(key)));
	}
}

export class ImmutableUint64Map extends MapUint64ToImmutableUint64 {
}

export class MapUint64ToMutableUint64 extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getUint64(key: u64): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.key(wasmtypes.uint64ToBytes(key)));
	}
}

export class MutableUint64Map extends MapUint64ToMutableUint64 {
}
