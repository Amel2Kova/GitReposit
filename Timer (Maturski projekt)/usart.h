#include "stm32f10x.h"

//��������� ��� ������������� GPIOA � USART1
GPIO_InitTypeDef    GPIO_InitStruct;
USART_InitTypeDef    USART_InitStruct;

void Init_usart(void); //���������� ������� ������������� ���������
void Usart1_Send_symbol(uint8_t); //���������� ������� �������� �������
void Usart1_Send_String(char* str); //���������� ������� �������� ������

char* str;
uint8_t data;
