#include <iostream>
#include <math.h>
#include <string>

using namespace std;

struct Bloco {
  unsigned long long int valor;
  unsigned long long int indice;
  Bloco *prox;
  Bloco() {
    valor = 0;
    indice = 0;
    prox = NULL;
  }
};

struct MyArray {
  unsigned long long int tamanho;
  Bloco *p;
  Bloco *u;
  MyArray() {
    tamanho = 0;
    p = NULL;
    u = NULL;
  }
};

void adiciona(int valor, int indice, MyArray *array) {
  Bloco *aux = new Bloco();
  if (array->tamanho == 0) {
    aux->valor = valor;
    aux->indice = indice;
    aux->prox = NULL;
    array->p = aux;
    array->u = aux;
    array->tamanho++;
  } else {
    aux->valor = valor;
    aux->indice = indice;
    aux->prox = NULL;
    array->u->prox = aux;
    array->u = aux;
    array->tamanho++;
  }
}

void printArray(MyArray *array) {
  if (array->tamanho == 0) {
    cout << "Array vázio!" << endl;
    return;
  }
  // 8
  Bloco *aux = array->p;
  cout << "[";
  while (aux != NULL) {
    cout << aux->valor;
    if (array->tamanho > 1 && aux->indice != array->u->indice)
      cout << ", ";
    aux = aux->prox;
  }
  cout << "]" << endl;
}

void printElemento(int indice, MyArray *array) {
  if (array->tamanho == 0) {
    cout << "Array vázio!" << endl;
    return;
  }
  // 8
  Bloco *aux = array->p;
  while (aux != NULL) {
    if (aux->indice == indice) {
      cout << aux->valor << endl;
      return;
    }
    aux = aux->prox;
  }
}

int getValor(int indice, MyArray *array) {
  Bloco *aux = array->p;
  while (aux != NULL) {
    if (aux->indice == indice)
      break;
    aux = aux->prox;
  }
  return aux->valor;
}

bool valor(int indice, MyArray *array) {
  Bloco *aux = array->p;
  while (aux != NULL) {
    if (aux->indice == indice)
      return true;
    aux = aux->prox;
  }
  return false;
}

void troca(int indice1, int indice2, MyArray *array) {
  if (array->tamanho == 0) {
    cout << "Array vázio!" << endl;
    return;
  }

  Bloco *aux = new Bloco();
  Bloco *aux1 = array->p;
  Bloco *aux2 = array->p;
  while (aux1->indice != indice1 || aux2->indice != indice2) {
    if (aux1->indice != indice1) {
      aux1 = aux1->prox;
    }
    if (aux2->indice != indice2) {
      aux2 = aux2->prox;
    }
  }

  aux->valor = aux1->valor;
  aux1->valor = aux2->valor;
  aux2->valor = aux->valor;
}

long long int bin(string a) {
  long long int soma = 0;
  long long int p = 0;
  for (int i = a.size() - 1; i > -1; i--) {
    if (a[i] != '0') {
      soma += pow(2, p);
    }
    p++;
  }
  return soma;
}
long long int rec(unsigned long long int k, MyArray *array) {
  //cout << k << endl;
  /*if(k < 0){
    return -1;
  }*/
  if (valor(k, array)) {
    //cout << k << "->" << getValor(k, array) << endl;
    return getValor(k, array);
  } else {
    adiciona(rec(k - 1, array) + rec(k - 2, array), k, array);
    return getValor(k, array);
  }
}

int main() {

  int t;
  char a, b, c;
  string n, x;
  unsigned long long int k;
  // fflush(stdin);
  while (cin >> t) {
    while (t > 0) {
      cin >> n;
      MyArray *l = new MyArray();
      adiciona(1, 1, l);
      adiciona(1, 2, l);
      k = bin(n);
      k = rec(k, l);
      x = to_string(k);
      if (x.size() > 3) {
        a = x[x.size() - 1];
        b = x[x.size() - 2];
        c = x[x.size() - 3];
        x = a + b + c;
        cout << x << endl;
      } else if (x.size() == 3) {
        cout << x << endl;
      } else if (x.size() == 2) {
        cout << "0" << x << endl;
      } else {
        cout << "00" << x << endl;
      }
      t--;
    }
  }

  return 0;
}
