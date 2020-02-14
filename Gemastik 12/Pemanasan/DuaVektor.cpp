#include <bits/stdc++.h>

using namespace std;

int main() {
    int value;
    int x[10];
    int y[10];
    int i = 0;
    int j = 0;

    while(cin >> value) {
        x[i] = value;
        i++;
    }

    while(cin >> value) {
        y[j] = value;
        j++;
    }
    cout << x[0];
    cout << y[0];
}
