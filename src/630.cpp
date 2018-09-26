#include<iostream>
#include<vector>
#include<algorithm>
#include<queue>
using namespace std;

struct Less {
    bool operator()(vector<int>& a, vector<int>& b) {
        // order by duration
        return a[0] < b[0];
    }
};

class Solution {
public:
    int scheduleCourse(vector<vector<int>>& courses) {
        sort(courses.begin(), courses.end(), [](vector<int>& a, vector<int>& b) {
            return a[1] < b[1];
        });
        int now = 0;
        int count = 0;
        priority_queue<vector<int>, vector<vector<int>>, Less> taken;
        for(auto c : courses) {
            if(now + c[0] <= c[1]) {
                now += c[0];
                count ++;
                taken.push(c);
            } else if(taken.size() > 0) {
                auto max_c = taken.top();
                if(max_c[0] > c[0]) {
                    taken.pop();
                    taken.push(c);
                    now += c[0] - max_c[0];
                }
            }
        }
        return count;
    }
};

int main(){
    Solution s;
    int raw_input[][2] = {{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}};
    vector<vector<int>> input;
    for(auto p : raw_input) {
        vector<int> c;
        c.push_back(p[0]);
        c.push_back(p[1]);
        input.push_back(c);
    }
    cout << s.scheduleCourse(input) << endl;

    int raw_input2[][2] = {{100,2},{32,50}};
    input.clear();
    for(auto p : raw_input2) {
        vector<int> c;
        c.push_back(p[0]);
        c.push_back(p[1]);
        input.push_back(c);
    }
    cout << s.scheduleCourse(input) << endl;
    return 0;
}
