#   R e p e t i t i o n  
  
 ! [ 1 5 6 3 5 3 7 6 8 7 0 4 4 ] ( C : \ U s e r s \ D E L L \ A p p D a t a \ R o a m i n g \ T y p o r a \ t y p o r a - u s e r - i m a g e s \ 1 5 6 3 5 3 7 6 8 7 0 4 4 . p n g )  
  
 �1u 
  
 ` ` ` g o  
 p a c k a g e   r o u t e r s  
  
 i m p o r t   (  
 	 " r e p e t i t i o n / c o n t r o l l e r s "  
  
 	 " g i t h u b . c o m / a s t a x i e / b e e g o "  
 )  
  
 f u n c   i n i t ( )   {  
 	 b e e g o . R o u t e r ( " / " ,   & c o n t r o l l e r s . M a i n C o n t r o l l e r { } )  
 	 b e e g o . R o u t e r ( " / l o g i n " ,   & c o n t r o l l e r s . M a i n C o n t r o l l e r { } ,   " g e t : S h o w L o g i n ; p o s t : H a n d l e L o g i n " )  
 	 b e e g o . R o u t e r ( " / i n d e x " ,   & c o n t r o l l e r s . M a i n C o n t r o l l e r { } ,   " g e t : S h o w I n d e x " )  
 	 b e e g o . R o u t e r ( " / a d d A r t i c l e " ,   & c o n t r o l l e r s . M a i n C o n t r o l l e r { } ,   " g e t : S h o w A d d ; p o s t : H a n d l e A d d " )  
 	 b e e g o . R o u t e r ( " / c o n t e n t " ,   & c o n t r o l l e r s . M a i n C o n t r o l l e r { } ,   " g e t : S h o w C o n t e n t ; p o s t : H a n d l e C o n t e n t " )  
 	 b e e g o . R o u t e r ( " / u p d a t e " ,   & c o n t r o l l e r s . M a i n C o n t r o l l e r { } ,   " g e t : S h o w U p d a t e ; p o s t : H a n d l e U p d a t e " )  
 }  
 ` ` `  
  
 