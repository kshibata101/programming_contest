#include <iostream>
#include <vector>
#define eol '\n';
using namespace std;

class Node {
public:
  int id;
  int l;
  int r;
  int p;
  Node(): id(-1),l(-1),r(-1),p(-1){};
};

Node* tree;

int create_tree(int parent, vector<int> pre, vector<int> in) {
    if (pre.size() == 0 && in.size() == 0) {
      return -1;
    }

    int root_id = pre[0];

    Node* node = &tree[root_id];
    node->p = parent;
    
    // left
    vector<int> left_pre;
    vector<int> left_in;

    int i = 0;
    int size = in.size();
    for (; i < size; i++) {
      if (in[i] == root_id) {
        break;
      }
      left_pre.push_back(pre[i+1]);
      left_in.push_back(in[i]);
    }
    if (left_pre.size()) {
      node->l = create_tree(root_id, left_pre, left_in);
    }

    vector<int> right_pre;
    vector<int> right_in;

    i++;
    for (;i < size; i++) {
      right_pre.push_back(pre[i]);
      right_in.push_back(in[i]);
    }
    if (right_pre.size()) {
      node->r = create_tree(root_id, right_pre, right_in);
    }

    return root_id;
}

vector<int> postorder(int id, vector<int> post) {
  Node* node = &tree[id];
  if (node->l != -1) {
    post = postorder(node->l, post);
  }
  if (node->r != -1) {
    post = postorder(node->r, post);
  }
  post.push_back(id);

  return post;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);
  
  int n,i;
  cin >> n;

  tree = new Node[n+1];

  vector<int> preorder(n);
  vector<int> inorder(n);

  for (i = 0; i < n; i++) {
    cin >> preorder[i];
  }
  for (i = 0; i < n; i++) {
    cin >> inorder[i];
  }

  int root = create_tree(-1, preorder, inorder);

  // post order
  vector<int> post;
  post = postorder(root, post);

  // output
  cout << post[0];
  for (i = 1; i < n; i++) {
    cout << " " << post[i];
  }
  cout << eol;
  
  return 0;
}
