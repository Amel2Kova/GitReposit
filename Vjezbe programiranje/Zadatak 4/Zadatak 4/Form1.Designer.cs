﻿namespace Zadatak_4
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
            this.cmdLabel = new System.Windows.Forms.Label();
            this.SuspendLayout();
            // 
            // cmdLabel
            // 
            this.cmdLabel.AutoSize = true;
            this.cmdLabel.Location = new System.Drawing.Point(109, 117);
            this.cmdLabel.Name = "cmdLabel";
            this.cmdLabel.Size = new System.Drawing.Size(35, 13);
            this.cmdLabel.TabIndex = 0;
            this.cmdLabel.Text = "label1";
            this.cmdLabel.Click += new System.EventHandler(this.cmdLabel_Click);
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(284, 261);
            this.Controls.Add(this.cmdLabel);
            this.Name = "Form1";
            this.Text = "Form1";
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.Label cmdLabel;
    }
}

