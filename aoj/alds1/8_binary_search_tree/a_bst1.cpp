#include <iostream>
#include <vector>
#define eol '\n';
using namespace std;

class Node {
public:
  Node* p;
  Node* l;
  Node* r;
  int k;
  Node(): p(nullptr),l(nullptr),r(nullptr),k(-1){};
};

Node* root = nullptr;

void insert(Node* z) {
  Node* y = nullptr;
  Node* x = root;
  while (x != nullptr) {
    y = x;
    if (z->k < x->k) {
      x = x->l;
    } else {
      x = x->r;
    }
  }
  z->p = y;
  
  if (y == nullptr) {
    root = z;
  } else if (z->k < y->k) {
    y->l = z;
  } else {
    y->r = z;
  }
}

void inorder(Node* node) {
  if (node->l != nullptr) {
    inorder(node->l);
  }
  cout << " " << node->k;
  if (node->r != nullptr) {
    inorder(node->r);
  }
  if (node->p == nullptr) {
    cout << eol;
  }
}

void preorder(Node* node) {
  cout << " " << node->k;
  if (node->l != nullptr) {
    preorder(node->l);
  }
  if (node->r != nullptr) {
    preorder(node->r);
  }
  if (node->p == nullptr) {
    cout << eol;
  }
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int m,k;
  string cmd;
  cin >> m;
  for (int i = 0; i < m; i++) {
    cin >> cmd;
    if (cmd == "insert") {
      cin >> k;
      Node* node = new Node();
      node->k = k;
      insert(node);
    } else if (cmd == "print") {
      inorder(root);
      preorder(root);
    }
  }
  
  return 0;
}
