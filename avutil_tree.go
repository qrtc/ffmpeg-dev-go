// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/tree.h>

typedef int (*av_tree_cmp_func)(void *opaque, void *elem);
typedef int (*av_tree_enu_func)(void *opaque, void *elem);

*/
import "C"
import "unsafe"

// AVTreeNode
type AVTreeNode C.struct_AVTreeNode

// typedef int (*av_tree_cmp_func)(void *opaque, void *elem);
type AvTreeCmpFunc = C.av_tree_cmp_func

// typedef int (*av_tree_enu_func)(void *opaque, void *elem);
type AvTreeEnuFunc = C.av_tree_enu_func

// AvTreeNodeAlloc allocates an AVTreeNode.
func AvTreeNodeAlloc() *AVTreeNode {
	return (*AVTreeNode)(C.av_tree_node_alloc())
}

// AvTreeFind finds an element.
func AvTreeFind(root *AVTreeNode, key CVoidPointer,
	cmp AvTreeCmpFunc, next []unsafe.Pointer) unsafe.Pointer {
	if next != nil && len(next) < 2 {
		panic("next len < 2")
	}
	return (unsafe.Pointer)(C.av_tree_find((*C.struct_AVTreeNode)(root),
		VoidPointer(key), (C.av_tree_cmp_func)(cmp),
		(*unsafe.Pointer)(unsafe.Pointer(&next[0]))))
}

// AvTreeInsert inserts or removes an element.
func AvTreeInsert(rootp **AVTreeNode, key CVoidPointer,
	cmp AvTreeCmpFunc, next **AVTreeNode) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_tree_insert(
		(**C.struct_AVTreeNode)(unsafe.Pointer(rootp)),
		VoidPointer(key), (C.av_tree_cmp_func)(cmp),
		(**C.struct_AVTreeNode)(unsafe.Pointer(next))))
}

// AvTreeDestroy
func AvTreeDestroy(t *AVTreeNode) {
	C.av_tree_destroy((*C.struct_AVTreeNode)(t))
}

// AvTreeEnumerate applies enu(opaque, &elem) to all the elements in the tree in a given range.
func AvTreeEnumerate(t *AVTreeNode, opaque CVoidPointer,
	cmp AvTreeCmpFunc, enu AvTreeEnuFunc) {
	C.av_tree_enumerate((*C.struct_AVTreeNode)(t), VoidPointer(opaque),
		(C.av_tree_cmp_func)(cmp), (C.av_tree_enu_func)(enu))
}
