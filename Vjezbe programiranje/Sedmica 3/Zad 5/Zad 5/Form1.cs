using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace Zad_5
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            comboBox1.Items.Add("Programiranje");

            comboBox1.Items.Add("Elektrotehnika");
            comboBox1.Items.Add("Historija");
            comboBox1.Items.Add("Tjelesni");
        }

        private void comboBox1_SelectedIndexChanged(object sender, EventArgs e)
        {
            textBox1.Text = "Odabrali ste : " +comboBox1.Text;
        }
    }
}
