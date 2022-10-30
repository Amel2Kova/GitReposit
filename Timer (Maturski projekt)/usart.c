#include "usart.h"

//������� ������������� ���������
void Init_usart()
{
  //�������� ������������ GPIOA, USART1 � �������������� ������� AFIO
  RCC_APB2PeriphClockCmd((RCC_APB2Periph_USART1 | RCC_APB2Periph_GPIOA | RCC_APB2Periph_AFIO), ENABLE);

  //������������� ������ PA9 - USART1_Tx
  GPIO_InitStruct.GPIO_Pin = GPIO_Pin_9; //��������� ������ PA9
  GPIO_InitStruct.GPIO_Speed = GPIO_Speed_50MHz; //�������� ����� ������������
  GPIO_InitStruct.GPIO_Mode = GPIO_Mode_AF_PP; //����� �������������� �������, ����� Push-Pull
  GPIO_Init(GPIOA, &GPIO_InitStruct); //�������� ��������� ��������� � ��������� GPIO�

  //������������� ������ PA10 - USART1_Rx
  GPIO_InitStruct.GPIO_Pin = GPIO_Pin_10; //��������� ������ PA10
  GPIO_InitStruct.GPIO_Mode  = GPIO_Mode_IN_FLOATING; //Input floating
  GPIO_Init(GPIOA, &GPIO_InitStruct); //�������� ��������� ��������� � ��������� GPIO�


  //������������� USART1
  USART_InitStruct.USART_BaudRate = 115200; //�������� ������ 9600 ���
  USART_InitStruct.USART_WordLength = USART_WordLength_8b; //����� ����� 8 ���
  USART_InitStruct.USART_StopBits = USART_StopBits_1; //1 ����-���
  USART_InitStruct.USART_Parity = USART_Parity_No ; //��� �������� ��������
  USART_InitStruct.USART_HardwareFlowControl = USART_HardwareFlowControl_None; //��� ����������� ��������
  USART_InitStruct.USART_Mode = USART_Mode_Rx | USART_Mode_Tx; //������� ���������� � �������� USART1
  USART_Init(USART1, &USART_InitStruct); //�������� ��������� ��������� � ��������� USART1

  USART_Cmd(USART1, ENABLE); //�������� USART1
}

//������� �������� �������
void Usart1_Send_symbol(uint8_t data)
{
  while(!(USART1->SR & USART_SR_TC)); //��������� ��������� ����� TC - ���������� ���������� ��������
  USART1->DR = data; //���������� �������� � ������� ������ - �������� ������
}
