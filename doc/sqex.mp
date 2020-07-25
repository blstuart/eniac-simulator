input eniac

beginfig(1);
initunit((-1.5in,1in));
mp((0in,1in));
acc((2in,1in));
acc((3.5in,1in));
constunit((5in,1in));
% label.top(btex Constant Transmitter etex, accnum((5in,1in)));
label.top("6", accnum((2in,1in)));
label.top("7", accnum((3.5in,1in)));
draw (2.1in,3.2in)--(7.0in,3.2in);
label.lft(btex $\scriptstyle 1$ etex, (2.12in,3.2in));
% draw (2.1in,3.3in)--(6.5in,3.3in);
% label.lft(btex $\scriptstyle 2$ etex, (2.12in,3.3in));
draw (-1.5in,0.8in)--(7.0in,0.8in);
label.lft(btex $\scriptstyle 1-1$ etex, (-1.5in,0.8in));
draw (-1.5in,0.7in)--(7.0in,0.7in);
label.lft(btex $\scriptstyle 1-2$ etex, (-1.5in,0.7in));
draw (-1.5in,0.6in)--(7.0in,0.6in);
label.lft(btex $\scriptstyle 1-3$ etex, (-1.5in, 0.6in));
% draw (-1.5in,0.5in)--(7.0in,0.5in);
% label.lft(btex $\scriptstyle 1-4$ etex, (-1.5in,0.5in));

drawarrow initio((-1.5in,1in))--initio((-1.5in,1in))+(0,-0.2in);
draw initio((-1.5in,1in))+(0,-0.2in)--initio((-1.5in,1in))+(0,-0.3in);	% p i.io 1-1
drawarrow mpi((0,1in),3)+(0,-0.51in)--mpi((0,1in),3)+(0,-0.31in);
draw mpi((0,1in),3)--mpi((0,1in),3)+(0,-0.31in);	% p 1-2 p.Ci
drawarrow mpoa((0,1in),3)--mpoa((0,1in),3)+(0,-0.24in);
draw mpoa((0,1in),3)+(0,-0.24in)--mpoa((0,1in),3)+(0,-0.54in);	% p p.C1o 1-3
% label.top("9", mpd((0,1in),4,1));	% s p.d17s1 9
label.top("9", mpd((0,1in),5,1));	% s p.d16s1 9
label.top("9", mpd((0,1in),6,1));	% s p.d15s1 9
label.top("9", mpd((0,1in),7,1));	% s p.d14s1 9
label.top("1", mpd((0,1in),7,2));	% s p.d14s2 1
label.top("2", mpsc((0,1in),3));	% s p.cC 2
drawarrow accA((2in,1in))--accA((2in,1in))+(0,0.25in);
draw accA((2in,1in))+(0,0.25in)--accA((2in,1in))+(0,0.35in);	% p a6.A 1
drawarrow accalpha((2in,1in))+(0,0.35in)--accalpha((2in,1in))+(0,0.25in);
draw accalpha((2in,1in))--accalpha((2in,1in))+(0,0.25in);	% p 1 a6.α
drawarrow acci((2in,1in),5)+(0,-0.3in)--acci((2in,1in),5)+(0,-0.2in);
draw acci((2in,1in),5)--acci((2in,1in),5)+(0,-0.2in);	% p 1-1 a6.5i
drawarrow acco((2in, 1in),5)--acco((2in,1in),5)+(0,-0.2in);
draw acco((2in,1in),5)+(0,-0.2in)--acco((2in, 1in),5)+(0,-0.4in);	% p a6.5o 1-2
drawarrow acci((2in,1in),6)+(0,-0.5in)--acci((2in,1in),6)+(0,-0.2in);
draw acci((2in,1in),6)--acci((2in,1in),6)+(0,-0.2in);	% p 1-3 a6.6i
label.top(btex $\alpha$ etex, accop((2in,1in),5));	% s a6.op5 α
label.top("A", accop((2in,1in),6));	% s a6.op6 A
label.top("1", accrp((2in,1in),5));	% s a6.rp5 1
label.top("2", accrp((2in,1in),6));	% s a6.rp6 2
drawarrow accalpha((3.5in,1in))+(0,0.35in)--accalpha((3.5in,1in))+(0,0.25in);
draw accalpha((3.5in,1in))--accalpha((3.5in,1in))+(0,0.25in);		% p 1 a7.α
drawarrow acco((3.5in,1in),6)--acco((3.5in,1in),6)+(0,-0.2in);
draw acco((3.5in,1in),6)+(0,-0.2in)--acco((3.5in,1in),6)+(0,-0.3in);	% p a7.6o 1-1
drawarrow acci((3.5in,1in),5)+(0,-0.3in)--acci((3.5in,1in),5)+(0,-0.2in);
draw acci((3.5in,1in),5)--acci((3.5in,1in),5)+(0,-0.2in);	% p 1-1 a7.5i
drawarrow acci((3.5in,1in),6)+(0,-0.5in)--acci((3.5in,1in),6)+(0,-0.2in);
draw acci((3.5in,1in),6)--acci((3.5in,1in),6)+(0,-0.2in);	% p 1-3 a7.6i
label.top(btex $\alpha$ etex, accop((3.5in,1in),5));	% s a20.op5 α
label.top(btex $\alpha$ etex, accop((3.5in,1in),6));	% s a20.op6 α
label.top("1", accrp((3.5in,1in),5));	% s a20.rp5 1
label.top("2", accrp((3.5in,1in),6));	% s a20.rp6 2
drawarrow consdo((5in,1in))--consdo((5in,1in))+(0,0.25in);
draw consdo((5in,1in))+(0,0.25in)--consdo((5in,1in))+(0,0.35in);	% p c.o 1
drawarrow consi((5in,1in),26)+(0,-0.41in)--consi((5in,1in),26)+(0,-0.31in);
draw consi((5in,1in), 26)--consi((5in,1in), 26)+(0,-0.31in);	% p 1-1 c.26i
label.top(btex $\rm J_{\rm lr}$ etex, consop((5in,1in),26));	% s c.s26 Jlr
label.lft("1", consj((5in,1in)));	% s c.j1 1

label.top(btex $\alpha$ etex, accalpha((2in,1in))+(0,-0.2in));
label.top(btex $\beta$ etex, accbeta((2in,1in))+(0,-0.23in));
label.top(btex $\gamma$ etex, accgamma((2in,1in))+(0,-0.2in));
label.top(btex $\delta$ etex, accdelta((2in,1in))+(0,-0.2in));
label.top(btex $\epsilon$ etex, accepsilon((2in,1in))+(0,-0.2in));
label.top(btex A etex, accA((2in,1in))+(0,-0.2in));
label.top(btex S etex, accS((2in,1in))+(0,-0.2in));
label.top(btex $\alpha$ etex, accalpha((3.5in,1in))+(0,-0.2in));
label.top(btex $\beta$ etex, accbeta((3.5in,1in))+(0,-0.23in));
label.top(btex $\gamma$ etex, accgamma((3.5in,1in))+(0,-0.2in));
label.top(btex $\delta$ etex, accdelta((3.5in,1in))+(0,-0.2in));
label.top(btex $\epsilon$ etex, accepsilon((3.5in,1in))+(0,-0.2in));
label.top(btex A etex, accA((3.5in,1in))+(0,-0.2in));
label.top(btex S etex, accS((3.5in,1in))+(0,-0.2in));

% draw accbeta((3.5in,1in))--accbeta((3.5in,1in))+(0,0.45in);	% p 2 a7.β
% draw acco((3.5in,1in),5)--acco((3.5in,1in),5)+(0,-0.4in);	% p a20.5o 1-2
% draw acco((3.5in,1in),6)--acco((3.5in,1in),6)+(0,-0.5in);	% p a20.6o 1-3
% draw initpi((-1.5in,1in))+(0,-0.4in)--initpi((-1.5in,1in));	% p 1-2 i.pi
% draw initpo((-1.5in,1in))--initpo((-1.5in,1in))+(0,-0.3in);	% p i.po 1-1
% s pr.2 P
% s pr.3 P
% s pr.15 P
% s pr.16 P

endfig;

end;
