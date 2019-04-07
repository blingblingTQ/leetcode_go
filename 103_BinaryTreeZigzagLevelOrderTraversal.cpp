 struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };

class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int> > result;
        if(root == NULL) {
            return result;
        }
        queue<TreeNode *> q;
        q.push(root);
        int direction = 1;
        while(!q.empty()) {
            int levelSize = q.size();
            vector<int> level;
            for(int i = 0; i < levelSize; i++) {
                TreeNode *node = q.front();
                q.pop();
                level.push_back(node->val);
                
                if(node->left != NULL) {
                    q.push(node->left);
                }
                if(node->right != NULL) {
                    q.push(node->right);
                }
            }
            
            if(direction == -1) {
                direction = 1;
                reverse(level.begin(), level.end());
                result.push_back(level);
            } else {
                direction = -1;
                result.push_back(level);
            }
        }
        return result;
    }
};
