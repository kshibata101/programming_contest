#include <iostream>
#define eol '\n';
using namespace std;

class Node {
public:
  int id;
  int p;
  int s;
  int c;
  int d;
  int h;
  string t;
  int l;
  int r;
  Node():id(-1),p(-1),s(-1),c(-1),d(-1),h(0){};
};

Node* tree;

void down(int id, int depth) {
  Node *node = &tree[id];
  node->d = depth;
  
  if (node->l != -1) down(node->l, depth+1);
  if (node->r != -1) down(node->r, depth+1);
}

void up(int id, int height) {
  Node *node = &tree[id];
  node->h = max(node->h, height);
  if (node->p != -1) up(node->p, height+1);
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n,id,l,r;
  Node *node, *child;
  cin >> n;
  tree = new Node[n];

  for (int i = 0; i < n; i++) {
    cin >> id >> l >> r;

    node = &tree[id];
    node->id = id;
    node->l = l;
    node->r = r;
    
    // child
    int c = 0;
    if (l != -1) {
      c++;

      child = &tree[l];
      child->p = id;
      child->s = r;
    }
    if (r != -1) {
      c++;

      child = &tree[r];
      child->p = id;
      child->s = l;
    }
    node->c = c;
  }

  for (int i = 0; i < n; i++) {
    node = &tree[i];
    if (node->p == -1) {
      node->t = "root";
      down(i, 0);
    } else if (node->c == 0) {
      node->t = "leaf";
      up(i, 0);
    } else {
      node->t = "internal node";
    }
  }

  for (int i = 0; i < n; i++) {
    node = &tree[i];
    
    cout << "node " << node->id 
         << ": parent = " << node->p 
         << ", sibling = " << node->s 
         << ", degree = " << node->c
         << ", depth = " << node->d
         << ", height = " << node->h
         << ", " << node->t
         << eol;
  }

  return 0;
}

