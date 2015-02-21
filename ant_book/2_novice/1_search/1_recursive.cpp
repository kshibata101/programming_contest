#include <iostream>
using namespace std;

int fact(int n) {
  if (n == 0) return 1;
  return n * fact(n - 1);
}

int fib (int n) {
  if (n <= 1) return 1;
  return fib(n-1) + fib(n-2);
}

int main() {
  int res = fact(5);
  cout << "fact 5: " << res << endl;

  res = fib(10);
  cout << "fib 10: " << res << endl;

  return 0;
}
