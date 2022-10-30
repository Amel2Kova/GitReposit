using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace Zad_2
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button5_Click(object sender, EventArgs e)
        {
            Application.Exit();
        }
        int R;
        private void button1_Click(object sender, EventArgs e)
        {
            R=int.Parse(textBox1.Text);
            for (int i = 0; i < R; i++)
            {
                listBox1.Items.Add(i);
            
            }
        }

        private void button3_Click(object sender, EventArgs e)
        {
            R = int.Parse(textBox1.Text);
            for (int i = 0; i < R; i++)
            {
                if((i%2)==0)
                listBox1.Items.Add(i);

            }
        }

        private void button2_Click(object sender, EventArgs e)
        { R = int.Parse(textBox1.Text);
            do{
                listBox1.Items.Add(R%10);
                R = R / 10;
            }while(R!=0);
        }

        private void button4_Click(object sender, EventArgs e)
        {
            listBox1.Items.Clear();
        }

       
    }
}
