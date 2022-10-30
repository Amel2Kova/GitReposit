using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace Zad_1
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void ADD_Click(object sender, EventArgs e)
        {
            listBox1.Items.Add(textBox1.Text);
        }

        int R;
        private void Remove_Click(object sender, EventArgs e)
        {
           
            listBox1.Items.Remove(textBox1.Text);

        }

        private void Insert_Click(object sender, EventArgs e)
        {
            R=int.Parse(textBox2.Text);
            listBox1.Items.Insert(R,textBox1.Text);
            R = 0;
        }

        private void Count_Click(object sender, EventArgs e)
        {
            textBox1.Text = Convert.ToString(listBox1.Items.Count);
           
        }

        private void Clear_Click(object sender, EventArgs e)
        {
            listBox1.Items.Clear();

        }

        private void Exit_Click(object sender, EventArgs e)
        {
            Application.Exit();
        }

       
    }
}
