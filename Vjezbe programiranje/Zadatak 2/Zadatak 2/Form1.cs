using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace Zadatak_2
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void cmdSnimi_Click(object sender, EventArgs e)
        {
            MessageBox.Show(txtIme.Text + " " + txtPrezime.Text);
        }
    }
}
