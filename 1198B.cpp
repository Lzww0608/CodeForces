#include <vector>
#include <iostream>
#include <algorithm>
#include <stack>
#include <cmath>
#include <string>
using namespace std;

struct event {
	int type, p, x;
};

int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int n;
	cin >> n;
	vector<int> a(n, 0);

	for (int& x : a) 
		cin >> x;

	int q;
	cin >> q;
	vector<event> events(q);
	vector<int> last_occurance(n, -1);
	for (int i = 0; i < q; ++i) {
		auto& ev = events[i];
		cin >> ev.type;
		if (ev.type == 1) {
			cin >> ev.p >> ev.x;
			ev.p--;
			last_occurance[ev.p] = i;
		} else {
			cin >> ev.x;
		}
	}
	vector<int> suffix_max(q + 1, 0);
	for (int i = q - 1; i >= 0; i--)
		suffix_max[i] = max(suffix_max[i + 1], events[i].type == 2 ? events[i].x : 0);

	for (int i = 0; i < n; ++i) {
		int balance = a[i];
		if (last_occurance[i] > -1)
			balance = events[last_occurance[i]].x;
		balance = max(balance, suffix_max[last_occurance[i] + 1]);
		cout << balance << " ";
	}
}