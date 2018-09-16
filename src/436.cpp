#include<iostream>
#include<map>
#include<vector>
using namespace std;

struct Interval {
      int start;
      int end;
      Interval() : start(0), end(0) {}
      Interval(int s, int e) : start(s), end(e) {}
 };

class Solution {
public:
    vector<int> findRightInterval(vector<Interval>& intervals) {
        map<int, int> starts;
        for(int i = 0; i < intervals.size(); i++) {
            auto interval = intervals[i];
            starts[interval.start] = i;
        }

        vector<int> ret;
        for(int i = 0; i < intervals.size(); i++) {
            auto r = starts.lower_bound(intervals[i].end);
            if (r == starts.end()) {
                ret.push_back(-1);
            } else {
                ret.push_back(r->second);
            }
        }
        return ret;
    }
};

void print(vector<int>& ret) {
    for (int i = 0; i < ret.size(); i++) {
        cout << ret[i] << " ";
    }
    cout << endl;
}

int main(){
    Solution s;
    vector<Interval> input;
    input.push_back(Interval(1, 2));
    print(s.findRightInterval(input));

    input.clear();
    input.push_back(Interval(3, 4));
    input.push_back(Interval(2, 3));
    input.push_back(Interval(1, 2));
    print(s.findRightInterval(input));

    input.clear();
    input.push_back(Interval(1, 4));
    input.push_back(Interval(2, 3));
    input.push_back(Interval(3, 4));
    print(s.findRightInterval(input));
    return 0;
}
