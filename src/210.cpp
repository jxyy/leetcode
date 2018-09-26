#include<iostream>
#include<vector>
#include<list>
using namespace std;

class Solution {
public:
    vector<int> findOrder(int numCourses, vector<pair<int, int>>& prerequisites) {
        vector<vector<int>> neighbors(numCourses);
        vector<int> indegree(numCourses, 0);
        for(auto p : prerequisites) {
            indegree[p.first] ++;
            neighbors[p.second].push_back(p.first);
        }

        list<int> zero_nodes;
        for(int i = 0; i < numCourses; i++){
            if(indegree[i] == 0) {
                zero_nodes.push_back(i);
            }
        }
        vector<int> order;
        while(!zero_nodes.empty()) {
            int node = zero_nodes.front();
            order.push_back(node);
            for(auto adj : neighbors[node]) {
                if(--(indegree[adj]) == 0){
                    zero_nodes.push_back(adj);
                }
            }
            zero_nodes.pop_front();
        }

        for(auto i : indegree) {
            if(i > 0) {
                order.clear();
                break;
            }
        }
        return order;
    }
};

void print(vector<int>& results) {
    for(auto i : results) {
        cout << i << " ";
    }
    cout << endl;
}

int main(){
    Solution s;
    vector<pair<int, int>> edges;
    edges.push_back(pair<int, int>(1, 0));
    edges.push_back(pair<int, int>(2, 0));
    edges.push_back(pair<int, int>(3, 1));
    edges.push_back(pair<int, int>(3, 2));
    print(s.findOrder(4, edges));

    edges.clear();
    edges.push_back(pair<int, int>(1, 0));
    edges.push_back(pair<int, int>(1, 2));
    edges.push_back(pair<int, int>(0, 1));
    print(s.findOrder(3, edges));
    return 0;
}
