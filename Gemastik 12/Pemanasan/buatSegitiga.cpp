#include <bits/stdc++.h>

using namespace std;

int main() {
    int a, i, j;
    char sym;

    cin >> a;
    cin >> sym;

    for (i=0; i<a; i++) {
        for (j=0; j<a-i-1; j++){
            cout << " ";
        }

        for (j=0; j<=i; j++){
            cout << sym << " ";
        }
        cout << endl;
    }

}
