using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace Zad4
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            checkedListBox1.Items.Add("jeda", CheckState.Unchecked);
            checkedListBox1.Items.Add("dva", CheckState.Unchecked);
            checkedListBox1.Items.Add("tri", CheckState.Unchecked);
            checkedListBox1.Items.Add("cetiri", CheckState.Unchecked);
            checkedListBox1.Items.Add("pet", CheckState.Unchecked);
            checkedListBox1.Items.Add("sest", CheckState.Unchecked);
            checkedListBox1.Items.Add("sedam", CheckState.Unchecked);
            checkedListBox1.Items.Add("osam", CheckState.Unchecked);
            checkedListBox1.Items.Add("devet", CheckState.Unchecked);
        }

        private void button1_Click(object sender, EventArgs e)
        {
            foreach(string X in checkedListBox1.CheckedItems ){
            
            
            listBox1.Items.Add(X);
        
            }
            
        }
    }
}
