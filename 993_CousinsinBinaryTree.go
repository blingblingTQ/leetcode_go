package main

 struct TreeNode {
      int val;
      TreeNode *left;
      TreeNode *right;
      TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };
 
class Solution {
public:
    bool isCousins(TreeNode* root, int x, int y) {
        if (root == NULL) {
            return false;
        }
        std::queue<std::pair<TreeNode *, TreeNode*> > q;
        auto rootPair = make_pair(root, (TreeNode*)NULL);
        q.push(rootPair);
        
        while(!q.empty()) {
            int currLevelSize = q.size();
            TreeNode *xParent = NULL;
            TreeNode *yParent = NULL;
            for(int i = 0; i < currLevelSize; i++) {
                auto nodePair = q.front();
                q.pop();
                
                if(nodePair.first->left != NULL) {
                    auto childPair = make_pair(nodePair.first->left, nodePair.first);
                    q.push(childPair);
                }
                if(nodePair.first->right != NULL) {
                    auto childPair = make_pair(nodePair.first->right, nodePair.first);
                    q.push(childPair);
                }
                
                if(nodePair.first->val == x) {
                    xParent = nodePair.second;
                }
                if(nodePair.first->val == y) {
                    yParent = nodePair.second;
                }
            }
            if(xParent && yParent) {
                //当前层都找到，则比较parent
                return xParent != yParent;
            } else if(xParent || yParent){
                //当前层一个找到一个没找到
                return false;
            }
        }
        return false;
    }
};
