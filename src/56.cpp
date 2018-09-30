#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

// Definition for an interval.
struct Interval {
    int start;
    int end;
    Interval() : start(0), end(0) {}
    Interval(int s, int e) : start(s), end(e) {}
};

class Solution {
public:
    vector<Interval> merge(vector<Interval>& intervals) {
        sort(intervals.begin(), intervals.end(), [](Interval& i1, Interval& i2){return i1.start < i2.start;});
        vector<Interval> ret;
        for(auto i : intervals) {
            if (ret.size() > 0 && i.start <= ret[ret.size()-1].end) {
                ret[ret.size()-1].end = max(i.end, ret[ret.size()-1].end);
            } else {
                ret.push_back(i);
            }
        }
        return ret;
    }
};

void print(vector<Interval>& v) {
    for(auto i : v) {
        cout << "[" << i.start << "," << i.end << "] " ;
    }
    cout << endl;
}

int main(){
    Solution s;
    vector<Interval> input;
    input.push_back(Interval(1, 3));
    input.push_back(Interval(2, 6));
    input.push_back(Interval(8, 10));
    input.push_back(Interval(15, 18));
    print(s.merge(input));

    input.clear();
    input.push_back(Interval(1, 4));
    input.push_back(Interval(4, 5));
    print(s.merge(input));

    input.clear();
    input.push_back(Interval(1, 4));
    input.push_back(Interval(2, 3));
    print(s.merge(input));
    return 0;
}
