namespace Zadatak_3
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
            this.label1 = new System.Windows.Forms.Label();
            this.cmdPrikaziPoruku = new System.Windows.Forms.Button();
            this.SuspendLayout();
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Location = new System.Drawing.Point(116, 39);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(170, 13);
            this.label1.TabIndex = 0;
            this.label1.Text = "Kliknite na dugme da vidite poruku";
            // 
            // cmdPrikaziPoruku
            // 
            this.cmdPrikaziPoruku.Location = new System.Drawing.Point(142, 71);
            this.cmdPrikaziPoruku.Name = "cmdPrikaziPoruku";
            this.cmdPrikaziPoruku.Size = new System.Drawing.Size(101, 23);
            this.cmdPrikaziPoruku.TabIndex = 1;
            this.cmdPrikaziPoruku.Text = "&Prikaži poruku";
            this.cmdPrikaziPoruku.UseVisualStyleBackColor = true;
            this.cmdPrikaziPoruku.Click += new System.EventHandler(this.cmdPrikaziPoruku_Click);
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(424, 145);
            this.Controls.Add(this.cmdPrikaziPoruku);
            this.Controls.Add(this.label1);
            this.Name = "Form1";
            this.Text = "Upotreba kontrola button i label";
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.Label label1;
        private System.Windows.Forms.Button cmdPrikaziPoruku;
    }
}

