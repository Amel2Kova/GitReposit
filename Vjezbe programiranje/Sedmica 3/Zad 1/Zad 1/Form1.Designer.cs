namespace Zad_1
{
    partial class Form1
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.listBox1 = new System.Windows.Forms.ListBox();
            this.textBox1 = new System.Windows.Forms.TextBox();
            this.ADD = new System.Windows.Forms.Button();
            this.Remove = new System.Windows.Forms.Button();
            this.Count = new System.Windows.Forms.Button();
            this.Insert = new System.Windows.Forms.Button();
            this.Exit = new System.Windows.Forms.Button();
            this.Clear = new System.Windows.Forms.Button();
            this.textBox2 = new System.Windows.Forms.TextBox();
            this.SuspendLayout();
            // 
            // listBox1
            // 
            this.listBox1.FormattingEnabled = true;
            this.listBox1.Location = new System.Drawing.Point(51, 57);
            this.listBox1.Name = "listBox1";
            this.listBox1.Size = new System.Drawing.Size(166, 303);
            this.listBox1.TabIndex = 0;
            // 
            // textBox1
            // 
            this.textBox1.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(238)));
            this.textBox1.Location = new System.Drawing.Point(223, 57);
            this.textBox1.Name = "textBox1";
            this.textBox1.Size = new System.Drawing.Size(167, 26);
            this.textBox1.TabIndex = 1;
            // 
            // ADD
            // 
            this.ADD.Font = new System.Drawing.Font("Microsoft Office Preview Font", 14.25F);
            this.ADD.Location = new System.Drawing.Point(295, 115);
            this.ADD.Name = "ADD";
            this.ADD.Size = new System.Drawing.Size(95, 30);
            this.ADD.TabIndex = 2;
            this.ADD.Text = "ADD";
            this.ADD.UseVisualStyleBackColor = true;
            this.ADD.Click += new System.EventHandler(this.ADD_Click);
            // 
            // Remove
            // 
            this.Remove.Font = new System.Drawing.Font("Microsoft Office Preview Font", 14.25F);
            this.Remove.Location = new System.Drawing.Point(295, 151);
            this.Remove.Name = "Remove";
            this.Remove.Size = new System.Drawing.Size(95, 30);
            this.Remove.TabIndex = 3;
            this.Remove.Text = "Remove";
            this.Remove.UseVisualStyleBackColor = true;
            this.Remove.Click += new System.EventHandler(this.Remove_Click);
            // 
            // Count
            // 
            this.Count.Font = new System.Drawing.Font("Microsoft Office Preview Font", 14.25F);
            this.Count.Location = new System.Drawing.Point(295, 223);
            this.Count.Name = "Count";
            this.Count.Size = new System.Drawing.Size(95, 30);
            this.Count.TabIndex = 5;
            this.Count.Text = "Count";
            this.Count.UseVisualStyleBackColor = true;
            this.Count.Click += new System.EventHandler(this.Count_Click);
            // 
            // Insert
            // 
            this.Insert.Font = new System.Drawing.Font("Microsoft Office Preview Font", 14.25F);
            this.Insert.Location = new System.Drawing.Point(295, 187);
            this.Insert.Name = "Insert";
            this.Insert.Size = new System.Drawing.Size(95, 30);
            this.Insert.TabIndex = 4;
            this.Insert.Text = "Insert";
            this.Insert.UseVisualStyleBackColor = true;
            this.Insert.Click += new System.EventHandler(this.Insert_Click);
            // 
            // Exit
            // 
            this.Exit.Font = new System.Drawing.Font("Microsoft Office Preview Font", 14.25F);
            this.Exit.Location = new System.Drawing.Point(295, 330);
            this.Exit.Name = "Exit";
            this.Exit.Size = new System.Drawing.Size(95, 30);
            this.Exit.TabIndex = 7;
            this.Exit.Text = "Exit";
            this.Exit.UseVisualStyleBackColor = true;
            this.Exit.Click += new System.EventHandler(this.Exit_Click);
            // 
            // Clear
            // 
            this.Clear.Font = new System.Drawing.Font("Microsoft Office Preview Font", 14.25F);
            this.Clear.Location = new System.Drawing.Point(295, 259);
            this.Clear.Name = "Clear";
            this.Clear.Size = new System.Drawing.Size(95, 30);
            this.Clear.TabIndex = 6;
            this.Clear.Text = "Clear";
            this.Clear.UseVisualStyleBackColor = true;
            this.Clear.Click += new System.EventHandler(this.Clear_Click);
            // 
            // textBox2
            // 
            this.textBox2.Location = new System.Drawing.Point(240, 196);
            this.textBox2.Name = "textBox2";
            this.textBox2.Size = new System.Drawing.Size(49, 20);
            this.textBox2.TabIndex = 8;
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(474, 391);
            this.Controls.Add(this.textBox2);
            this.Controls.Add(this.Exit);
            this.Controls.Add(this.Clear);
            this.Controls.Add(this.Count);
            this.Controls.Add(this.Insert);
            this.Controls.Add(this.Remove);
            this.Controls.Add(this.ADD);
            this.Controls.Add(this.textBox1);
            this.Controls.Add(this.listBox1);
            this.Name = "Form1";
            this.Text = "Form1";
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.ListBox listBox1;
        private System.Windows.Forms.TextBox textBox1;
        private System.Windows.Forms.Button ADD;
        private System.Windows.Forms.Button Remove;
        private System.Windows.Forms.Button Count;
        private System.Windows.Forms.Button Insert;
        private System.Windows.Forms.Button Exit;
        private System.Windows.Forms.Button Clear;
        private System.Windows.Forms.TextBox textBox2;
    }
}

