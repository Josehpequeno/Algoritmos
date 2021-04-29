#include <iostream>
//#include <boost/utility/binary.hpp>
using namespace std;
//g++ - lboost_regex boost.cpp - o boost
struct Bloco
{
  int valor;
  int indice;
  Bloco *prox;
  Bloco()
  {
    valor = 0;
    indice = 0;
    prox = NULL;
  }
};

struct MyArray
{
  int tamanho;
  Bloco *p;
  Bloco *u;
  MyArray()
  {
    tamanho = 0;
    p = NULL;
    u = NULL;
  }
};

void adiciona(int valor, MyArray *array)
{
  Bloco *aux = new Bloco();
  if (array->tamanho == 0)
  {
    aux->valor = valor;
    aux->indice = array->tamanho;
    aux->prox = NULL;
    array->p = aux;
    array->u = aux;
    array->tamanho++;
  }
  else
  {
    aux->valor = valor;
    aux->indice = array->tamanho;
    aux->prox = NULL;
    array->u->prox = aux;
    array->u = aux;
    array->tamanho++;
  }
}

void printArray(MyArray *array)
{
  if (array->tamanho == 0)
  {
    cout << "Array vázio!" << endl;
    return;
  }
  // 8
  Bloco *aux = array->p;
  cout << "[";
  while (aux != NULL)
  {
    cout << aux->valor;
    if (array->tamanho > 1 && aux->indice != array->u->indice)
      cout << ", ";
    aux = aux->prox;
  }
  cout << "]" << endl;
}

void printElemento(int indice, MyArray *array)
{
  if (array->tamanho == 0)
  {
    cout << "Array vázio!" << endl;
    return;
  }
  // 8
  Bloco *aux = array->p;
  while (aux != NULL)
  {
    if (aux->indice == indice)
    {
      cout << aux->valor << endl;
      return;
    }
    aux = aux->prox;
  }
}

int getValor(int indice, MyArray *array)
{
  Bloco *aux = array->p;
  while (aux != NULL)
  {
    if (aux->indice == indice)
      break;
    aux = aux->prox;
  }
  return aux->valor;
}

void troca(int indice1, int indice2, MyArray *array)
{
  if (array->tamanho == 0)
  {
    cout << "Array vázio!" << endl;
    return;
  }

  Bloco *aux = new Bloco();
  Bloco *aux1 = array->p;
  Bloco *aux2 = array->p;
  while (aux1->indice != indice1 || aux2->indice != indice2)
  {
    if (aux1->indice != indice1)
    {
      aux1 = aux1->prox;
    }
    if (aux2->indice != indice2)
    {
      aux2 = aux2->prox;
    }
  }

  aux->valor = aux1->valor;
  aux1->valor = aux2->valor;
  aux2->valor = aux->valor;
}

void insertSort(MyArray *l)
{
  for (int j = 1; j < l->tamanho; j++)
  {
    int chave = getValor(j, l);
    int i = j - 1;
    while (i >= 0 && getValor(i, l) > chave)
    {
      troca(i + 1, i, l);
      //l[i + 1] = l[i];
      i = i - 1;
    }
    //l[i + 1] = chave;
  }
}

int main()
{

  MyArray *l = new MyArray();
  int t;
  char a, b, c;
  string n, x;
  while (cin >> t)
  {
    while (t > 0)
    {
      cin >> n;
      unsigned short n1 = BOOST_BINARY(n);
      cout << n1 << endl;
      /*x = rec(n, lista);
      if (len(x) > 3)
      {
        a = x[l->tamanho - 1];
        b = x[l->tamanho - 2];
        c = x[l->tamanho - 3];
        x = a << b << c;
        cout << x << endl;
      }
      else if (len(x) == 3)
      {
        cout << x << endl;
      }
      else if (len(x) == 2)
      {
        cout << "0" << x << endl;
      }
      else
      {
        cout << "00" << x << endl;
      }*/
      t--;
    }

    //fflush(stdin);
    while (teste != '\n')
    {
      cin >> a;
      adiciona(a, l);
      //teste = getchar(); //para capturar o '\n'
    }
  }

  return 0;
}
