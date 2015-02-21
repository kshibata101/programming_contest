#include <iostream>
#include <stack>

using namespace std;

int cnt;

void insertion_sort(int* A, int n, int g) {
  for (int i = g; i < n; i++) {
    int key = A[i];
    int j = i - g;
    while (j >= 0 && A[j] > key) {
      A[j+g] = A[j];
      j -= g;
      cnt++;
    }
    A[j+g] = key;
  }
}

void shell_sort (int* A, int n) {
  cnt = 0;

  int m = 0;
  int val = 1;
  stack<int> st;
  while(val <= n) {
    m++;
    st.push(val);

    val = 3 * val + 1;
  }

  // line 1
  cout << m << endl;

  for (int i = 0; i < m; i++) {
    int g = st.top();
    st.pop();

    // line 2
    if (i == 0) {
      cout << g;
    } else {
      cout << " " << g;
    }

    insertion_sort(A, n, g);
  }
  cout << endl << cnt << endl;
  
  for (int i = 0; i < n; i++) {
    cout << A[i] << endl;
  }
}

int main() {
  int n;
  int* A;

  cin >> n;
  A = new int[n];
  for (int i = 0; i < n; i++) {
    cin >> A[i];
  }

  shell_sort(A, n);

  return 0;
}

