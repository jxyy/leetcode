#include<iostream>
#include<vector>
#include<unordered_map>
using namespace std;


// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> neighbors;

    Node() {
        // cout << "new one!" << endl;
    }

    Node(int _val, vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};


class Solution {
public:
    Node* cloneGraph(Node* node) {
        if(node == NULL) return NULL;

        unordered_map<Node*, Node*> map;
        vector<Node*> pending;
        pending.push_back(node);
        Node* newg = new Node();
        newg->val = node->val;
        map[node] = newg;
        while (!pending.empty()) {
            auto old_node = pending[pending.size()-1];
            auto new_node = map[old_node];
            pending.pop_back();
            for(auto neighbor : old_node->neighbors) {
                if(map.find(neighbor) == map.end()) {
                    cout << "new neighbor" << endl;
                    auto new_neighbor = new Node();
                    new_neighbor->val = neighbor->val;
                    map[neighbor] = new_neighbor;
                    pending.push_back(neighbor);
                    new_node->neighbors.push_back(new_neighbor);
                } else {
                    auto new_neighbor = map[neighbor];
                    new_node->neighbors.push_back(new_neighbor);
                }
                cout << "new edge" << endl;
            }
            
        }
        return newg;
    }
};

int main(){
    Node* g1 = new Node();
    g1->val = 1;
    Node* g2 = new Node();
    g2->val = 2;
    Node* g3 = new Node();
    g3->val = 3;
    Node* g4 = new Node();
    g4->val = 4;

    g1->neighbors.push_back(g2);
    g1->neighbors.push_back(g4);

    g2->neighbors.push_back(g1);
    g2->neighbors.push_back(g3);

    g3->neighbors.push_back(g2);
    g3->neighbors.push_back(g4);

    g4->neighbors.push_back(g1);
    g4->neighbors.push_back(g3);

    Solution s;
    s.cloneGraph(g1);
    return 0;
}
