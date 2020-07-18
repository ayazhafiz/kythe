//- @a defines/binding A
//- A code ACode
//- ACode child.0 AContext
//- AContext.kind "CONTEXT"
//- AContext.pre_text "function"
//- ACode child.1 ASpace
//- ASpace.kind "BOX"
//- ASpace.pre_text " "
//- ACode child.2 AName
//- AName.kind "IDENTIFIER"
//- AName.pre_text "a"
//- ACode child.3 AParams
//- AParams.kind "PARAMETER_LOOKUP_BY_PARAM"
//- AParams.pre_text "("
//- AParams.post_child_text ", "
//- AParams.post_text ")"
//- ACode child.4 ATy
//- ATy.kind "TYPE"
//- ATy.pre_text ": "
//- ATy.post_text "Function"
function a(): Function {
  //- @#0b defines/binding B
  //- B code BCode
  //- BCode child.0 BContext
  //- BContext.pre_text "(local function)"
  //- BCode child.2 BName
  //- BName.pre_text "b"
  //- BCode child.3 _BParams
  //- BCode child.4 BTy
  //- BTy.post_text "string | number"
  function b(): string|number {
    return 1;
  }
  return b;
}

//- @counter defines/binding Counter
//- Counter code CounterCode
//- CounterCode child.0 CounterContext
//- CounterContext.pre_text "function"
//- CounterCode child.2 CounterName
//- CounterName.pre_text "counter"
//- CounterCode child.3 _CounterParams
//- CounterCode child.4 CounterTy
//- CounterTy.post_text "Generator<number, { iter: number; }, unknown>"
function* counter() {
  let i = 0;
  for (;;) {
    if (yield++ i) {
      return {iter: i};
    }
  }
}
