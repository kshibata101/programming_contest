#include <iostream>
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

void preorder(int id) {
  Node *node = &tree[id];
  
  cout << " " << node->id;
  
  if (node->l != -1) {
    preorder(node->l);
  }
  if (node->r != -1) {
    preorder(node->r);
  }
}

void inorder(int id) {
  Node *node = &tree[id];
 
  if (node->l != -1) {
    inorder(node->l);
  }

  cout << " " << node->id;

  if (node->r != -1) {
    inorder(node->r);
  }
}

void postorder(int id) {
  Node *node = &tree[id];
 
  if (node->l != -1) {
    postorder(node->l);
  }

  if (node->r != -1) {
    postorder(node->r);
  }

  cout << " " << node->id;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n,id,l,r;
  Node *node;

  cin >> n;
  tree = new Node[n];
  for (int i = 0; i < n; i++) {
    cin >> id >> l >> r;
    
    node = &tree[id];
    node->id = id;
    node->l  = l;
    node->r  = r;
    if (l != -1) {
      node = &tree[l];
      node->p = id;
    }
    if (r != -1) {
      node = &tree[r];
      node->p = id;
    }
  }

  int parent_id = -1;
  for (int i = 0; i < n; i++) {
    node = &tree[i];
    if (node->p == -1) {
      parent_id = i;
      break;
    }
  }

  cout << "Preorder" << eol;
  preorder(parent_id);
  cout << eol;
  cout << "Inorder" << eol;
  inorder(parent_id);
  cout << eol;
  cout << "Postorder" << eol;
  postorder(parent_id);
  cout << eol;
  
  return 0;
}
