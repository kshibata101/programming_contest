#include <iostream>

using namespace std;

typedef struct {
  int id;
  int parent;
  int depth;
  int child_num;
  int* children;
} node;

node* tree;

void update_depth(int node_id, int depth) {
  node nod = tree[node_id];
  tree[node_id].depth = depth;
  int child_num = nod.child_num;
  for (int i = 0; i < child_num; i++) {
    update_depth(nod.children[i], depth+1);
  }
}

int main() {
  int n, id, k, i,j;
  int* children;
  node nod,child;

  cin >> n;

  tree = new node[n];

  for (i = 0; i < n; i++) {
    tree[i].parent = -1;
  }

  for (i = 0; i < n; i++) {
    cin >> id >> k;

    nod = tree[id];
    nod.id = id;
    nod.child_num = k;

    children = new int[k];
    for (j = 0; j < k; j++) {
      cin >> children[j];

      tree[children[j]].parent = id;
    }

    nod.children = children;

    tree[id] = nod;
  }

  // depth
  for (i = 0; i < n; i++) {
    nod = tree[i];
    if (nod.parent == -1) {
      update_depth(i,0);
    }
  }

  // output
  for (i = 0; i < n; i++) {
    nod = tree[i];
    cout << "node " << nod.id << ": parent = " << nod.parent <<  ", depth = " << nod.depth << ", ";
    if (nod.parent == -1) {
      cout << "root";
    } else if (nod.child_num == 0) {
      cout << "leaf";
    } else {
      cout << "internal node";
    }
    cout << ", [";
    if (nod.child_num != 0) {
      cout << nod.children[0];
      for (j = 1; j < nod.child_num; j++) {
        cout << ", " << nod.children[j];
      }
    }
    cout << "]" << endl;
  }
  
  return 0;
}
