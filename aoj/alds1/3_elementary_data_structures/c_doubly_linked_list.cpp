#include <iostream>
#include <list>
#include <sstream>

using namespace std;

list<int> xlist;
int x;

void insert(int x) {
  xlist.push_front(x);
}

void deleteValue(int x) {
  for (list<int>::iterator itr = xlist.begin(); itr != xlist.end(); ++itr) {
    if (*itr == x) {
      xlist.erase(itr);
      break;
    }
  }
}

void deleteFirst() {
  xlist.pop_front();
}

void deleteLast() {
  xlist.pop_back();
}

int main() {
  ios::sync_with_stdio(false);
  
  int n;
  cin >> n, cin.ignore();
  
  string line, cmd;
  
  for (int i = 0; i < n; i++) {
    getline(cin, line);
    stringstream ss(line);
    ss >> cmd;
    
    if (cmd == "insert") {
      ss >> x;
      insert(x);
    } else if (cmd == "delete") {
      ss >> x;
      deleteValue(x);
    } else if (cmd == "deleteFirst") {
      deleteFirst();
    } else if (cmd == "deleteLast") {
      deleteLast();
    }
  }

  // cout
  for (list<int>::iterator itr = xlist.begin(); itr != xlist.end(); ++itr) {
    if (itr != xlist.begin()) {
      cout << " ";
    }
    cout << *itr;
  }
  cout << endl;
  return 0;
}

