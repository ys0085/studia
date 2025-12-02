with Ada.Text_IO; use Ada.Text_IO;
with Ada.Numerics.Float_Random; use Ada.Numerics.Float_Random;
with Random_Seeds; use Random_Seeds;
with Ada.Real_Time; use Ada.Real_Time;

procedure  Travelers is


   -- Travelers moving on the board

   Nr_Of_Travelers : constant Integer :=15;

   Min_Steps : constant Integer := 10 ;
   Max_Steps : constant Integer := 100 ;

   Min_Delay : constant Duration := 0.01;
   Max_Delay : constant Duration := 0.05;
   
   Max_Move_Attempts : constant Integer := 10; 
   Move_Timeout : constant Duration := 0.4; 

   -- 2D Board with torus topology

   Board_Width  : constant Integer := 15;
   Board_Height : constant Integer := 15;

   -- Timing

   Start_Time : Time := Clock;  -- global startnig time

   -- Random seeds for the tasks' random number generators
   
   Seeds : Seed_Array_Type(1..Nr_Of_Travelers) := Make_Seeds(Nr_Of_Travelers);

   -- Types, procedures and functions

   -- Postitions on the board
   type Position_Type is record	
      X: Integer range 0 .. Board_Width - 1; 
      Y: Integer range 0 .. Board_Height - 1; 
   end record;	   

      type Board_Array_Type is array (Integer range <>, Integer range <>) of Boolean;
   -- Protected Board to track occupied positions and ensure synchronization
   protected type Protected_Board_Type is
      function Is_Occupied(X, Y: Integer) return Boolean;
      procedure Mark_Occupied(X, Y: Integer; Occupied: Boolean);
      procedure Move_If_Possible(From_X, From_Y, To_X, To_Y: Integer; Success: out Boolean);
   private
      Board: Board_Array_Type(0 .. Board_Width - 1, 0 .. Board_Height - 1) := (others => (others => False));
   end Protected_Board_Type;
   
   protected body Protected_Board_Type is
      function Is_Occupied(X, Y: Integer) return Boolean is
      begin
         return Board(X, Y);
      end Is_Occupied;
      
      procedure Mark_Occupied(X, Y: Integer; Occupied: Boolean) is
      begin
         Board(X, Y) := Occupied;
      end Mark_Occupied;
      
      procedure Move_If_Possible(From_X, From_Y, To_X, To_Y: Integer; Success: out Boolean) is
      begin
         if not Is_Occupied (To_X, To_Y) then
            Mark_Occupied(From_X, From_Y, False);
            Mark_Occupied(To_X, To_Y, True);
            Success := True;
         else
            Success := False;
         end if;
      end Move_If_Possible;
   end Protected_Board_Type;
   
   Game_Board : Protected_Board_Type;

   -- elementary steps
   procedure Move_Down( Position: in out Position_Type ) is
   begin
      Position.Y := ( Position.Y + 1 ) mod Board_Height;
   end Move_Down;

   procedure Move_Up( Position: in out Position_Type ) is
   begin
      Position.Y := ( Position.Y + Board_Height - 1 ) mod Board_Height;
   end Move_Up;

   procedure Move_Right( Position: in out Position_Type ) is
   begin
      Position.X := ( Position.X + 1 ) mod Board_Width;
   end Move_Right;

   procedure Move_Left( Position: in out Position_Type ) is
   begin
      Position.X := ( Position.X + Board_Width - 1 ) mod Board_Width;
   end Move_Left;

   -- traces of travelers
   type Trace_Type is record 	      
      Time_Stamp:  Duration;	      
      Id : Integer;
      Position: Position_Type;      
      Symbol: Character;	      
   end record;	      

   type Trace_Array_type is  array(0 .. Max_Steps) of Trace_Type;

   type Traces_Sequence_Type is record
      Last: Integer := -1;
      Trace_Array: Trace_Array_type ;
   end record; 


   procedure Print_Trace( Trace : Trace_Type ) is
      Symbol : String := ( ' ', Trace.Symbol );
   begin
      Put_Line(
         Duration'Image( Trace.Time_Stamp ) & " " &
         Integer'Image( Trace.Id ) & " " &
         Integer'Image( Trace.Position.X ) & " " &
         Integer'Image( Trace.Position.Y ) & " " &
         ( ' ', Trace.Symbol ) -- print as string to avoid: '
         );
   end Print_Trace;

   procedure Print_Traces( Traces : Traces_Sequence_Type ) is
   begin
      for I in 0 .. Traces.Last loop
         Print_Trace( Traces.Trace_Array( I ) );
      end loop;
   end Print_Traces;

   -- task Printer collects and prints reports of traces
   task Printer is
      entry Report( Traces : Traces_Sequence_Type );
   end Printer;
   
   task body Printer is 
   begin
      for I in 1 .. Nr_Of_Travelers loop -- range for TESTS !!!
         accept Report( Traces : Traces_Sequence_Type ) do
            Print_Traces( Traces );
         end Report;
         end loop;
   end Printer;


   -- travelers
   type Traveler_Type is record
      Id: Integer;
      Symbol: Character;
      Position: Position_Type;    
   end record;


   task type Traveler_Task_Type is	
      entry Init(Id: Integer; Seed: Integer; Symbol: Character);
      entry Start;
   end Traveler_Task_Type;	

   task body Traveler_Task_Type is
      G : Generator;
      Traveler : Traveler_Type;
      Time_Stamp : Duration;
      Nr_of_Steps: Integer;
      Traces: Traces_Sequence_Type; 
      Is_Trapped : Boolean := False;

      procedure Store_Trace is
      begin  
         Traces.Last := Traces.Last + 1;
         Traces.Trace_Array( Traces.Last ) := ( 
            Time_Stamp => Time_Stamp,
            Id => Traveler.Id,
            Position => Traveler.Position,
            Symbol => Traveler.Symbol
         );
      end Store_Trace;
      
      procedure Make_Step is
         N : Integer; 
         New_Position : Position_Type;
         Valid_Move : Boolean := False;
         Start_Attempt_Time : Time;
         Attempts : Integer := 0;
      begin
         Valid_Move := False;
         Start_Attempt_Time := Clock;

         while not Valid_Move and then 
               Attempts < Max_Move_Attempts and then
               To_Duration(Clock - Start_Attempt_Time) < Move_Timeout loop
               
            N := Integer( Float'Floor(2.0 * Random(G)) );
            Attempts := Attempts + 1;

            if Traveler.Id mod 2 = 0 then
               N := N + 2;
            end if;
            
            -- Prepare potential new position
            New_Position := Traveler.Position;
            
            case N is
               when 0 =>
                  Move_Up(New_Position);
               when 1 =>
                  Move_Down(New_Position);
               when 2 =>
                  Move_Left(New_Position);
               when 3 =>
                  Move_Right(New_Position);
               when others =>
                  Put_Line( " ?????????????? " & Integer'Image( N ) );
               end case;
            
            
            Game_Board.Move_If_Possible(Traveler.Position.X, Traveler.Position.Y, 
                                                New_Position.X, New_Position.Y, Valid_Move);
            
            if Valid_Move then
               Traveler.Position := New_Position;
            end if;
            delay 0.005;
         end loop;
         
         
         if not Valid_Move then
            Is_Trapped := True;
            if Character'Pos(Traveler.Symbol) < Character'Pos('a') then
               Traveler.Symbol := Character'Val(Character'Pos(Traveler.Symbol) + 32); -- Convert to lowercase
            end if;
         end if;
      end Make_Step;

   begin
      accept Init(Id: Integer; Seed: Integer; Symbol: Character) do
         Reset(G, Seed); 
         Traveler.Id := Id;
         Traveler.Symbol := Symbol;
         declare
         Init_Attempts : Integer := 0;
         begin
            loop
               Init_Attempts := Init_Attempts + 1;
               Traveler.Position := (
                  X => Id,
                  Y => Id           
                  );
                  
               exit when not Game_Board.Is_Occupied(Traveler.Position.X, Traveler.Position.Y) or else
                           Init_Attempts >= 100;
                           
               delay 0.005;
            end loop;
            Game_Board.Mark_Occupied(Traveler.Position.X, Traveler.Position.Y, True);
         end;
         
         Store_Trace; -- store starting position
         -- Number of steps to be made by the traveler  
         Nr_of_Steps := Min_Steps + Integer( Float(Max_Steps - Min_Steps) * Random(G));
         -- Time_Stamp of initialization
         Time_Stamp := To_Duration ( Clock - Start_Time ); -- reads global clock
      end Init;
      
      -- wait for initialisations of the remaining tasks:
      accept Start do
         null;
      end Start;

      for Step in 0 .. Nr_of_Steps loop
         delay Min_Delay+(Max_Delay-Min_Delay)*Duration(Random(G));
         -- do action ...
         if not Is_Trapped then
            Make_Step;
         end if;
         Store_Trace;
         Time_Stamp := To_Duration ( Clock - Start_Time ); -- reads global clock
      end loop;
      Printer.Report( Traces );
   end Traveler_Task_Type;


   -- local for main task

   Travel_Tasks: array (0 .. Nr_Of_Travelers-1) of Traveler_Task_Type; -- for tests
   Symbol : Character := 'A';
   begin 
   
   -- Prit the line with the parameters needed for display script:
   Put_Line(
         "-1 "&
         Integer'Image( Nr_Of_Travelers ) &" "&
         Integer'Image( Board_Width ) &" "&
         Integer'Image( Board_Height )      
      );


   -- init tarvelers tasks
   for I in Travel_Tasks'Range loop
      Travel_Tasks(I).Init( I, Seeds(I+1), Symbol );   -- `Seeds(I+1)` is ugly :-(
      Symbol := Character'Succ( Symbol );
   end loop;

   -- start tarvelers tasks
   for I in Travel_Tasks'Range loop
      Travel_Tasks(I).Start;
   end loop;

end Travelers;