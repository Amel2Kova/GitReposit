using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace Zadatak_2_Sesija_2
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            string odabir="Odabrano je :  ";
            if (checkBox1.Checked)
                odabir = odabir + checkBox1.Text+"  ";

            if (checkBox2.Checked)
                odabir = odabir + checkBox2.Text + "  ";

            if (checkBox3.Checked)
                odabir = odabir + checkBox3.Text + "  ";
            MessageBox.Show(odabir);
        }
    }
}
