#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int* make_array(string s) {
  int* arr = new int[26];
  for (int i = 0; i < 26; i++) {
    arr[i] = 0;
  }
  for (int i = 0; i < s.length(); i++) {
    int key = s[i] - 'a';
    arr[key]++;
  }
  return arr;
}

bool same(int* arr1, int* arr2) {
  for (int i = 0; i < 26; i++) {
    if (arr1[i] != arr2[i]) {
      return false;
    }
  }
  return true;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n;
  cin >> n;

  int* indeed = make_array("indeednow");
  
  string s;
  for (int i = 0; i < n; i++) {
    cin >> s;
    int* arr = make_array(s);
    if (same(indeed, arr)) {
      cout << "YES" << eol;
    } else {
      cout << "NO" << eol;
    }
  }
  
  return 0;
}
