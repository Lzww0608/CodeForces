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

	int T, n;
	cin >> T;
	while (T--) {
		cin >> n;
		vector<int> a(10);
		int x = 0;
		for (int i = 0; i < n; i++) {
			cin >> x;
			a[x % 10]++;
		}
		for (int i = 0; i < 10; i++) {
			if (a[i] == 0) continue;
			a[i]--;
			for (int j = i; j < 10; j++) {
				if (a[j] == 0) continue;
				a[j]--;
				for (int k = j; k < 10; k++) {
					if (a[k] > 0 && (k + j + i) % 10 == 3) {
						cout << "YES" << endl;
						goto next;
					}
				}
				a[j]++;
			}
			a[i]++;
		}
		cout << "NO" << endl;
	next:;
	}

	return 0;
}