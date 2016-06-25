package main

import (
	"fmt"
)

func main() {
    ip := ip_range{}
    ip.init(1, 3, "rus")
    fmt.Println(ip)
    ip1 := ip_range{}
    ip1.init(4, 7, "eng")
    fmt.Println(ip1)
    fmt.Println(ip1.gn(ip))
    arr := []ip_range{ip_range{8,10, "rus"}, ip_range{1, 2, "eng"}, ip_range{19,21, "fr"}, ip_range{3,4, "us"}, ip_range{5,7, "aus"}}
    bin := binary_heap{}
    fmt.Println(arr)
    bin.built_heap(arr)
    fmt.Println(bin.get_range_key(8))
}

type ip_range struct {
    start int
    back int
    key string
}

func (ip *ip_range) init(start int, back int, key string) {
    ip.start = start
    ip.back = back
    ip.key = key
}

func (ip *ip_range) gn(other ip_range) bool {
    if ip.start > other.back {
        return true
    }
    return false
}

func (ip *ip_range) eq(other ip_range) bool {
    if ip.start == other.start && ip.back == other.back {
        return true
    }
    return false
}

func (ip *ip_range) in(val int) bool {
    fmt.Println("===", ip.start, ip.back, val)
    if ip.start <= val && val <= ip.back {
        return true
    }
    return false
}

func (ip *ip_range) go_to_child(val int) bool {
    if ip.back <= val {
        return true
    }
    return false
}

type binary_heap struct {
	heap []ip_range
}

func (h *binary_heap) size() int {
	return len(h.heap) - 1
}

func (h *binary_heap) heapify(root int) {
	for true {
		left_ch := 2*root + 1
		right_ch := 2*root + 2
		largest := root

		if left_ch <= h.size() && h.heap[left_ch].gn(h.heap[largest]) {
			largest = left_ch
		}

		if right_ch <= h.size() && h.heap[right_ch].gn(h.heap[largest]) {
			largest = right_ch
		}

		if largest == root {
			break
		}

		it := h.heap[largest]
		h.heap[largest] = h.heap[root]
		h.heap[root] = it
		root = largest
	}
}

func (h *binary_heap) get_range_key(val int) string {
    root := 0
    if h.size() != 0 && h.heap[root].in(val) {
        return h.heap[root].key
    }
    for true {
        left_ch := 2*root + 1
		right_ch := 2*root + 2
        fmt.Println(left_ch, right_ch, h.size())
        if left_ch > h.size() || right_ch > h.size() {
            break
        }

        if h.heap[left_ch].in(val) {
            return h.heap[left_ch].key
        }

        if h.heap[right_ch].in(val) {
            return h.heap[right_ch].key
        }

        if h.heap[root].go_to_child(val) {
            root = right_ch
        } else {
            root = left_ch
        }
    }
    return "not key"

}

func (h *binary_heap) built_heap(data []ip_range) {
	h.heap = data
	num := h.size() / 2
	for i := num; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *binary_heap) get_max() ip_range {
	res := h.heap[0]
	h.heap[0] = h.heap[h.size()]
	h.heap = h.heap[0:h.size()]
	h.heapify(0)
	return res
}

func (h *binary_heap) heap_sort(data []ip_range) []ip_range {
	h.built_heap(data)
	res := make([]ip_range, len(data))
	num := h.size()
	for i := 0; i <= num; i++ {

		res[i] = h.get_max()
	}
	return res
}
