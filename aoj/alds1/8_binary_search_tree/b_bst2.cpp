#include <iostream>
#define eol '\n';
using namespace std;

class Node {
public:
  int k;
  Node* p;
  Node* l;
  Node* r;
  Node(): k(-1),p(nullptr),l(nullptr),r(nullptr){};
};

Node* root;

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

bool find(int k) {
  Node* x = root;
  while (x != nullptr) {
    if (k == x->k) {
      return true;
    } else if (k < x->k) {
      x = x->l;
    } else {
      x = x->r;
    }
  }
  return false;
}

void inorder(Node* node) {
  if (node == nullptr) {
    return;
  }
  
  inorder(node->l);
  cout << " " << node->k;
  inorder(node->r);
  if (node->p == nullptr) {
    cout << eol;
  }
}

void preorder(Node* node) {
  if (node == nullptr) {
    return;
  }
  cout << " " << node->k;
  preorder(node->l);
  preorder(node->r);
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
    } else if (cmd == "find") {
      cin >> k;
      if (find(k)) {
        cout << "yes" << eol;
      } else {
        cout << "no" << eol;
      }
    } else if (cmd == "print") {
      inorder(root);
      preorder(root);
    }
  }

  return 0;
}
