#include <iostream>
#include <vector>
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

void insert(int k) {
  Node* y = nullptr;
  Node* x = root;
  while (x != nullptr) {
    y = x;
    if (k < x->k) {
      x = x->l;
    } else {
      x = x->r;
    }
  }

  Node* z = new Node();
  z->k = k;
  z->p = y;

  if (y == nullptr) {
    root = z;
  } else if (k < y->k){
    y->l = z;
  } else {
    y->r = z;
  }
}

bool find(int k) {
  Node* x = root;
  while (x != nullptr) {
    if (x->k == k) {
      return true;
    } else if(k < x->k) {
      x = x->l;
    } else {
      x = x->r;
    }
  }
  return false;
}

void remove(int k) {
  Node* z = nullptr;
  Node* x = root;
  while (x != nullptr) {
    if (x->k == k) {
      z = x;
      break;
    } else if(k < x->k) {
      x = x->l;
    } else {
      x = x->r;
    }
  }
  
  if (z->l == nullptr && z->r == nullptr) {
    // has not child
    if (z->p->l != nullptr && z->p->l->k == z->k) {
      z->p->l = nullptr;
    } else {
      z->p->r = nullptr;
    }
  } else if (z->l == nullptr) {
    // has one child
    if (z->p->l != nullptr && z->p->l->k == z->k) {
      z->p->l = z->r;
      z->r->p = z->p;
    } else {
      z->p->r = z->r;
      z->r->p = z->p;
    }
  } else if (z->r == nullptr) {
    if (z->p->l != nullptr && z->p->l->k == z->k) {
      z->p->l = z->l;
      z->l->p = z->p;
    } else {
      z->p->r = z->l;
      z->l->p = z->p;
    }
  } else {
    Node* w = z->r;
    while (w->l != nullptr) {
      w = w->l;
    }

    // まずwを取り除く
    remove(w->k);

    // 次にzの親の子をwに更新する
    if (z->p->l != nullptr && z->p->l->k == z->k) {
      z->p->l = w;
    } else {
      z->p->r = w;
    }
    // wの親をzの親に
    w->p = z->p;
    // wの子をzの子に
    w->l = z->l;
    w->r = z->r;
    // zの子の親をwに(ただし、w自身の場合は更新せず)
    z->l->p = w;
    if (z->r != nullptr)
      z->r->p = w;
  }
}

void inorder(Node* x) {
  if (x == nullptr) return;
    
  inorder(x->l);
  cout << " " << x->k;
  inorder(x->r);
  if (x->k == root->k)
    cout << eol;
}

void preorder(Node* x) {
  if (x == nullptr) return;
  cout << " " << x->k;
  preorder(x->l);
  preorder(x->r);
  if (x->k == root->k)
    cout << eol;
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
      insert(k);
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
    } else if (cmd == "delete") {
      cin >> k;
      remove(k);
    }
  }

  return 0;
}
