#include<iostream>
#include<vector>
using namespace std;

class Iterator {
    struct Data;
	Data* data;
public:
	Iterator(const vector<int>& nums);
	Iterator(const Iterator& iter);
	virtual ~Iterator();
	// Returns the next element in the iteration.
	int next();
	// Returns true if the iteration has more elements.
	bool hasNext() const;
};


class PeekingIterator : public Iterator {
    bool is_cached;
    int cached_num;
public:
	PeekingIterator(const vector<int>& nums) : Iterator(nums) {
	    // Initialize any member here.
	    // **DO NOT** save a copy of nums and manipulate it directly.
	    // You should only use the Iterator interface methods.
	    is_cached = false;
	}

    // Returns the next element in the iteration without advancing the iterator.
	int peek() {
        if (!is_cached) {
            cached_num = Iterator::next();
            is_cached = true;
        }
        return cached_num;
	}

	// hasNext() and next() should behave the same as in the Iterator interface.
	// Override them if needed.
	int next() {
	    if(is_cached) {
            is_cached = false;
            return cached_num;
        }
        return Iterator::next();
	}

	bool hasNext() const {
	    return is_cached || Iterator::hasNext();
	}
};

int main(){
    return 0;
}
