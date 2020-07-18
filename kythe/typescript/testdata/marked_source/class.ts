interface I {
  imethod(): string|void;
}

abstract class C implements I {
  //- @constructor defines/binding CCtor
  //- CCtor code CCtorCode
  //- CCtorCode child.0 CCtorContext
  //- CCtorContext.pre_text "constructor"
  //- CCtorCode child.2 CCtorName
  //- CCtorName.pre_text "C"
  //- CCtorCode child.3 _CCtorParams
  //- CCtorCode child.4 CCtorTy
  //- CCtorTy.post_text "C"
  constructor() {}
  //- @imethod defines/binding Cimethod
  //- Cimethod code CimethodCode
  //- CimethodCode child.0 CimethodContext
  //- CimethodContext.pre_text "(method)"
  //- CimethodCode child.2 CimethodName
  //- CimethodName.pre_text "C.imethod"
  //- CimethodCode child.3 _CimethodParams
  //- CimethodCode child.4 CimethodTy
  //- CimethodTy.post_text "void"
  imethod() {}
  //- @m1 defines/binding Cm1
  //- Cm1 code Cm1Code
  //- Cm1Code child.0 Cm1Context
  //- Cm1Context.pre_text "(method)"
  //- Cm1Code child.2 Cm1Name
  //- Cm1Name.pre_text "C.m1"
  //- Cm1Code child.3 _Cm1Params
  //- Cm1Code child.4 Cm1Ty
  //- Cm1Ty.post_text "void"
  m1() {}
  //- @m2 defines/binding Cm2
  //- Cm2 code Cm2Code
  //- Cm2Code child.0 Cm2Context
  //- Cm2Context.pre_text "(method)"
  //- Cm2Code child.2 Cm2Name
  //- Cm2Name.pre_text "C.m2"
  //- Cm2Code child.3 _Cm2Params
  //- Cm2Code child.4 Cm2Ty
  //- Cm2Ty.post_text "void"
  static m2() {}
  //- @m3 defines/binding Cm3
  //- Cm3 code Cm3Code
  //- Cm3Code child.0 Cm3Context
  //- Cm3Context.pre_text "(method)"
  //- Cm3Code child.2 Cm3Name
  //- Cm3Name.pre_text "C.m3<T = number>"
  //- Cm3Code child.3 _Cm3Params
  //- Cm3Code child.4 Cm3Ty
  //- Cm3Ty.post_text "T"
  abstract m3<T = number>(): T;
  //- @m4 defines/binding Cm4
  //- Cm4 code Cm4Code
  //- Cm4Code child.0 Cm4Context
  //- Cm4Context.pre_text "(method)"
  //- Cm4Code child.2 Cm4Name
  //- Cm4Name.pre_text "C.m4"
  //- Cm4Code child.3 _Cm4Params
  //- Cm4Code child.4 Cm4Ty
  //- Cm4Ty.post_text "Promise<void>"
  async m4() {}
  //- @m5 defines/binding Cm5
  //- Cm5 code Cm5Code
  //- Cm5Code child.0 Cm5Context
  //- Cm5Context.pre_text "(method)"
  //- Cm5Code child.2 Cm5Name
  //- Cm5Name.pre_text "C.m5"
  //- Cm5Code child.3 _Cm5Params
  //- Cm5Code child.4 Cm5Ty
  //- Cm5Ty.post_text "void"
  private m5() {}
  //- @m6 defines/binding Cm6
  //- Cm6 code Cm6Code
  //- Cm6Code child.0 Cm6Context
  //- Cm6Context.pre_text "(method)"
  //- Cm6Code child.2 Cm6Name
  //- Cm6Name.pre_text "C.m6"
  //- Cm6Code child.3 _Cm6Params
  //- Cm6Code child.4 Cm6Ty
  //- Cm6Ty.post_text "void"
  protected m6() {}
  //- @m7 defines/binding Cm7
  //- Cm7 code Cm7Code
  //- Cm7Code child.0 Cm7Context
  //- Cm7Context.pre_text "(method)"
  //- Cm7Code child.2 Cm7Name
  //- Cm7Name.pre_text "C.m7"
  //- Cm7Code child.3 _Cm7Params
  //- Cm7Code child.4 Cm7Ty
  //- Cm7Ty.post_text "any"
  m7?();
  //- @mem defines/binding CmemGetter
  //- CmemGetter code CmemGetterCode
  //- CmemGetterCode child.0 CmemGetterContext
  //- CmemGetterContext.pre_text "(getter)"
  //- CmemGetterCode child.2 CmemGetterName
  //- CmemGetterName.pre_text "C.mem"
  //- CmemGetterCode child.3 CmemGetterTy
  //- CmemGetterTy.post_text "number"
  get mem() {
    return 1;
  }
  //- @mem defines/binding CmemSetter
  //- CmemSetter code CmemSetterCode
  //- CmemSetterCode child.0 CmemSetterContext
  //- CmemSetterContext.pre_text "(setter)"
  //- CmemSetterCode child.2 CmemSetterName
  //- CmemSetterName.pre_text "C.mem"
  //- CmemSetterCode child.3 CmemSetterTy
  //- CmemSetterTy.post_text "number"
  set mem(newMem) {}
}

const g = class {
  // TODO(ayazhafiz): this isn't picked up
  constructor() {}
  //- @foo defines/binding Gfoo
  //- Gfoo code GfooCode
  //- GfooCode child.0 GfooContext
  //- GfooContext.pre_text "(method)"
  //- GfooCode child.2 GfooName
  //- GfooName.pre_text "g.foo"
  //- GfooCode child.3 _GfooParams
  //- GfooCode child.4 GfooTy
  //- GfooTy.post_text "void"
  foo() {}
}
