#include <vector>
#include <iostream>
#include <algorithm>
#include <cmath>
using namespace std;

int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int n, m;
	cin >> n >> m;
	vector<int> a(n, 0);
	long long x, y1 = INT_MAX, y2 = INT_MIN;
	for (auto& num : a) {
		cin >> num;
	}
	for (int i = 0; i < m; ++i) {
		cin >> x;
		y1 = min(x, y1);
		y2 = max(x, y2);
	}
	sort(a.begin(), a.end());
	long long x1 = a[0], x2 = a.back();
	long long ans1 = max(x1 * y1, x1 * y2), ans2 = max(x2 * y1, x2 * y2);
	if (ans1 > ans2) {
		x1 = a[1];
	} else {
		x2 = a[a.size() - 2];
	}
	cout << max(max(x1 * y1, x1 * y2), max(x2 * y1, x2 * y2)) << endl;
	return 0;
}
