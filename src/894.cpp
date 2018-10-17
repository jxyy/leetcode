#include<iostream>
#include<vector>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
 
class Solution {
    TreeNode* genTree(vector<int> v) {
        if(v.size() == 0){
            return NULL;
        }
        TreeNode *node = new TreeNode(0);
        return node;
    }
public:
    vector<TreeNode*> allPossibleFBT(int N) {
        
    }
};

int main(){
    return 0;
}
