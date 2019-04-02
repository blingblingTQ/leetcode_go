class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};

class Solution {
public:
    
    //dfs
    void dfs(Node *root, vector<vector<int> >& result, int level) {
        if(root == NULL) {
            return;
        }
        
        if(level == result.size()) {
            vector<int> v;
            result.push_back(v);
        }
        
        result[level].push_back(root->val);
        for(int i = 0; i < root->children.size(); i++) {
            dfs(root->children[i], result, level+1);
        }
    }
    
    vector<vector<int>> levelOrder(Node* root) {
        vector<vector<int> > result;
        dfs(root, result, 0);
        return result;
    }


    //bfs
    vector<vector<int>> levelOrder(Node* root) {
        vector<vector<int> > result;
        if(root == NULL) {
            return result;
        }
        queue<Node *> q;
        q.push(root);

        while(!q.empty()) {
            int size = q.size();
            vector<int> currLevel(size);
            for(int i = 0; i < size; i++) {
                Node *node = q.front();
                q.pop();
                currLevel[i] = node->val;

                for(int j = 0; j < node->children.size(); j++) {
                    q.push(node->children[j]);
                }
            }
            result.push_back(currLevel);
        }
        return result;
    }
};
