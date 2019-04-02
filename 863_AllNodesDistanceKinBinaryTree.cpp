 struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };
 
class Solution {
public:
    
    void dfs(TreeNode *root, TreeNode *parent, unordered_map<TreeNode*, TreeNode*>& parents) {
        if(root == NULL) {
            return;
        }
        parents[root] = parent;
        dfs(root->left, root, parents);
        dfs(root->right, root, parents);
    }
    
    vector<int> distanceK(TreeNode* root, TreeNode* target, int K) {
        vector<int> result;
        if(root == NULL) {
            return result;
        }
        unordered_map<TreeNode*, TreeNode*> parents;
        unordered_map<TreeNode *, bool> seen;
        dfs(root, nullptr, parents);
        int dist = 0;
        
        queue<TreeNode *> q;
        q.push(target);
        
        while(!q.empty()) {
            int size = q.size();
            for(int i = 0; i < size; i++) {
                TreeNode *node = q.front();
                q.pop();
                seen[node] = true;
                if(dist == K) {
                    result.push_back(node->val);
                }
                
                if(node->left && !seen[node->left]) {
                    q.push(node->left);
                }
                if(node->right && !seen[node->right]) {
                    q.push(node->right);
                }
                if(parents[node] && !seen[parents[node]]) {
                    q.push(parents[node]);
                }
            }
            dist++;
        }
        return result;
    }
};
