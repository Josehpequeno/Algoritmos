#include <iostream>
using namespace std;

void insertSort(int (&l)[6]){//array por referencia
  int size = sizeof(l)/4;
  for (int j = 1; j < size;j++){
    int chave = l[j];
    int i = j-1;
    while (i>=0 && l[i]>chave)
    {
      l[i+1] = l[i];
      i = i -1;
    }
    l[i+1] = chave;    
  }
}

int main() {

  int l[] = {5,2,4,6,1,3};
  
  for (int i; i < 6;i++){
    cout << l[i] << ' ';
  }
  cout << "\n";
  insertSort(l);
  for (int i; i < 6;i++){
    cout << l[i] << ' ';
  }
  cout << "\n";
  return 0;
}
