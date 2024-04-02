#include <iostream>
#include <vector>
#include <unordered_map>
#include <cstdlib>
#include <cmath>
#include <algorithm>
#include <string>
using namespace std;
using ll = long long;


int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	cout.tie(nullptr);

	int  n;
	cin >> n;
	vector<int> a(n), pre(n + 1);
	pre[0] = INT_MAX;
	for (int i = 0; i < n; i++) {
		cin >> a[i];
		pre[i+1] = pre[i] & ~a[i];
	}
	int res = 0, idx = 0, suf = INT_MAX;
	for (int i = n - 1; i >= 0; i--) {
		int t = a[i] & pre[i] & suf;
		if (t >= res) {
			res = t;
			idx = i;
		}
		suf &= ~a[i];
	}
	cout << a[idx] << " ";
	for (int i = 0; i < n; i++) {
		if (i != idx)
			cout << a[i] << " ";
	}

	return 0;
}