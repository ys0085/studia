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

-- 2D Board with torus topology

  Board_Width  : constant Integer := 5;
  Board_Height : constant Integer := 5;

-- Timing

  Start_Time : Time := Clock;  -- global startnig time
  
-- ADDED: Maximum time a traveler can wait before timing out
  Max_Wait_Time : constant Duration := 2.0;

-- Random seeds for the tasks' random number generators
 
  Seeds : Seed_Array_Type(1..Nr_Of_Travelers) := Make_Seeds(Nr_Of_Travelers);

-- Types, procedures and functions

  -- Postitions on the board
  type Position_Type is record	
    X: Integer range 0 .. Board_Width - 1; 
    Y: Integer range 0 .. Board_Height - 1; 
  end record;	   
  
  -- ADDED: Define array type for board occupancy
  type Occupancy_Array is array(0 .. Board_Width - 1, 0 .. Board_Height - 1) of Boolean;

  -- ADDED: Protected object for board occupancy management
  protected Board_Manager is
    procedure Occupy(Pos: Position_Type; Success: out Boolean);
    procedure Release(Pos: Position_Type);
    function Is_Occupied(Pos: Position_Type) return Boolean;
  private
    Board: Occupancy_Array := (others => (others => False));
  end Board_Manager;

  protected body Board_Manager is
    procedure Occupy(Pos: Position_Type; Success: out Boolean) is
    begin
      if Board(Pos.X, Pos.Y) then
        Success := False;
      else
        Board(Pos.X, Pos.Y) := True;
        Success := True;
      end if;
    end Occupy;

    procedure Release(Pos: Position_Type) is
    begin
      Board(Pos.X, Pos.Y) := False;
    end Release;

    function Is_Occupied(Pos: Position_Type) return Boolean is
    begin
      return Board(Pos.X, Pos.Y);
    end Is_Occupied;
  end Board_Manager;

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
    -- ADDED: Flag for timed out status
    Timed_Out: Boolean := False;
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
    -- ADDED: Variables for timeout tracking
    Move_Success: Boolean;
    Last_Move_Time: Time;
    Wait_Time: Duration;
    Original_Symbol: Character;
    Timeout_Count: Natural := 0;

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
    
    -- MODIFIED: Now returns success status and manages board occupancy
    function Make_Step return Boolean is
      N : Integer;
      New_Position: Position_Type;
      Try_Count: Integer := 0;
      Max_Tries: constant Integer := 4; -- Try all four directions
      Directions: array(0..3) of Integer := (0, 1, 2, 3);
      Dir_Index: Integer;
      Success: Boolean := False;
    begin
      -- Try each direction in random order until one works or all fail
      while Try_Count < Max_Tries and not Success loop
        -- Pick a random direction from remaining options
        N := Integer(Float'Floor(Float(Max_Tries - Try_Count) * Random(G)));
        Dir_Index := N;
        
        -- Make a copy of current position to try move
        New_Position := Traveler.Position;
        
        case Directions(Dir_Index) is
          when 0 =>
            Move_Up(New_Position);
          when 1 =>
            Move_Down(New_Position);
          when 2 =>
            Move_Left(New_Position);
          when 3 =>
            Move_Right(New_Position);
          when others =>
            Put_Line(" ?????????????? " & Integer'Image(N));
        end case;
        
        -- Try to occupy the new position
        Board_Manager.Occupy(New_Position, Success);
        
        if Success then
          -- Release old position only after successfully occupying new position
          Board_Manager.Release(Traveler.Position);
          -- Update traveler position
          Traveler.Position := New_Position;
          return True;
        end if;
        
        -- Swap tried direction with last untried direction
        Directions(Dir_Index) := Directions(Max_Tries - Try_Count - 1);
        Directions(Max_Tries - Try_Count - 1) := Directions(Dir_Index);
        Try_Count := Try_Count + 1;
      end loop;
      
      return False; -- No move was possible
    end Make_Step;

    -- ADDED: Randomize the travelers initial position until an empty spot is found
    procedure Find_Empty_Initial_Position is
      Found: Boolean := False;
    begin
      while not Found loop
        Traveler.Position := (
          X => Integer(Float'Floor(Float(Board_Width) * Random(G))),
          Y => Integer(Float'Floor(Float(Board_Height) * Random(G)))
        );
        Board_Manager.Occupy(Traveler.Position, Found);
      end loop;
    end Find_Empty_Initial_Position;

  begin
    accept Init(Id: Integer; Seed: Integer; Symbol: Character) do
      Reset(G, Seed); 
      Traveler.Id := Id;
      Traveler.Symbol := Symbol;
      Original_Symbol := Symbol;
      
      
      Find_Empty_Initial_Position;
      
      
      Store_Trace; -- store starting position
      -- Number of steps to be made by the traveler  
      Nr_of_Steps := Min_Steps + Integer(Float(Max_Steps - Min_Steps) * Random(G));
      -- Time_Stamp of initialization
      Time_Stamp := To_Duration(Clock - Start_Time); -- reads global clock
      Last_Move_Time := Clock; -- ADDED: Initialize last move time
    end Init;
    
    -- wait for initialisations of the remaining tasks:
    accept Start do
      null;
    end Start;

    for Step in 0 .. Nr_of_Steps loop
      delay Min_Delay+(Max_Delay-Min_Delay)*Duration(Random(G));
      
      
      Move_Success := Make_Step;
      
      -- ADDED: Timeout handling
      if Move_Success then
        Last_Move_Time := Clock;
        -- If traveler was timed out and now can move, restore uppercase symbol
        if Traveler.Timed_Out then
          Traveler.Symbol := Original_Symbol;
          Traveler.Timed_Out := False;
        end if;
      else
        -- Check if waited too long
        Wait_Time := To_Duration(Clock - Last_Move_Time);
        if Wait_Time > Max_Wait_Time and not Traveler.Timed_Out then
          -- Time out the traveler
          Traveler.Symbol := Character'Val(Character'Pos(Traveler.Symbol) + 32); -- Convert to lowercase
          Traveler.Timed_Out := True;
          Timeout_Count := Timeout_Count + 1;
        end if;
      end if;
      
      Store_Trace;
      Time_Stamp := To_Duration(Clock - Start_Time); -- reads global clock
    end loop;
    
    -- ADDED: Release position before terminating
    Board_Manager.Release(Traveler.Position);
    
    Printer.Report(Traces);
  end Traveler_Task_Type;


-- local for main task

  Travel_Tasks: array (0 .. Nr_Of_Travelers-1) of Traveler_Task_Type; -- for tests
  Symbol : Character := 'A';
begin 
  
  -- Prit the line with the parameters needed for display script:
  Put_Line(
      "-1 "&
      Integer'Image(Nr_Of_Travelers) &" "&
      Integer'Image(Board_Width) &" "&
      Integer'Image(Board_Height)      
    );

  -- init tarvelers tasks
  for I in Travel_Tasks'Range loop
    Travel_Tasks(I).Init(I, Seeds(I+1), Symbol);   -- `Seeds(I+1)` is ugly :-(
    Symbol := Character'Succ(Symbol);
  end loop;



  -- start tarvelers tasks
  for I in Travel_Tasks'Range loop
    Travel_Tasks(I).Start;
  end loop;

end Travelers;