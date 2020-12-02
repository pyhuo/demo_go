.
├── Make.dist
├── README.vendor
├── all.bash
├── all.bat
├── all.rc
├── archive
│   ├── tar
│   │   ├── common.go
│   │   ├── example_test.go
│   │   ├── format.go
│   │   ├── reader.go
│   │   ├── reader_test.go
│   │   ├── stat_actime1.go
│   │   ├── stat_actime2.go
│   │   ├── stat_unix.go
│   │   ├── strconv.go
│   │   ├── strconv_test.go
│   │   ├── tar_test.go
│   │   ├── testdata
│   │   │   ├── file-and-dir.tar
│   │   │   ├── gnu-incremental.tar
│   │   │   ├── gnu-long-nul.tar
│   │   │   ├── gnu-multi-hdrs.tar
│   │   │   ├── gnu-nil-sparse-data.tar
│   │   │   ├── gnu-nil-sparse-hole.tar
│   │   │   ├── gnu-not-utf8.tar
│   │   │   ├── gnu-sparse-big.tar
│   │   │   ├── gnu-utf8.tar
│   │   │   ├── gnu.tar
│   │   │   ├── hardlink.tar
│   │   │   ├── hdr-only.tar
│   │   │   ├── invalid-go17.tar
│   │   │   ├── issue10968.tar
│   │   │   ├── issue11169.tar
│   │   │   ├── issue12435.tar
│   │   │   ├── neg-size.tar
│   │   │   ├── nil-uid.tar
│   │   │   ├── pax-bad-hdr-file.tar
│   │   │   ├── pax-bad-mtime-file.tar
│   │   │   ├── pax-global-records.tar
│   │   │   ├── pax-multi-hdrs.tar
│   │   │   ├── pax-nil-sparse-data.tar
│   │   │   ├── pax-nil-sparse-hole.tar
│   │   │   ├── pax-nul-path.tar
│   │   │   ├── pax-nul-xattrs.tar
│   │   │   ├── pax-path-hdr.tar
│   │   │   ├── pax-pos-size-file.tar
│   │   │   ├── pax-records.tar
│   │   │   ├── pax-sparse-big.tar
│   │   │   ├── pax.tar
│   │   │   ├── small.txt
│   │   │   ├── small2.txt
│   │   │   ├── sparse-formats.tar
│   │   │   ├── star.tar
│   │   │   ├── trailing-slash.tar
│   │   │   ├── ustar-file-devs.tar
│   │   │   ├── ustar-file-reg.tar
│   │   │   ├── ustar.tar
│   │   │   ├── v7.tar
│   │   │   ├── writer-big-long.tar
│   │   │   ├── writer-big.tar
│   │   │   ├── writer.tar
│   │   │   └── xattrs.tar
│   │   ├── writer.go
│   │   └── writer_test.go
│   └── zip
│       ├── example_test.go
│       ├── reader.go
│       ├── reader_test.go
│       ├── register.go
│       ├── struct.go
│       ├── testdata
│       │   ├── crc32-not-streamed.zip
│       │   ├── dd.zip
│       │   ├── go-no-datadesc-sig.zip
│       │   ├── go-with-datadesc-sig.zip
│       │   ├── gophercolor16x16.png
│       │   ├── readme.notzip
│       │   ├── readme.zip
│       │   ├── symlink.zip
│       │   ├── test-trailing-junk.zip
│       │   ├── test.zip
│       │   ├── time-22738.zip
│       │   ├── time-7zip.zip
│       │   ├── time-go.zip
│       │   ├── time-infozip.zip
│       │   ├── time-osx.zip
│       │   ├── time-win7.zip
│       │   ├── time-winrar.zip
│       │   ├── time-winzip.zip
│       │   ├── unix.zip
│       │   ├── utf8-7zip.zip
│       │   ├── utf8-infozip.zip
│       │   ├── utf8-osx.zip
│       │   ├── utf8-winrar.zip
│       │   ├── utf8-winzip.zip
│       │   ├── winxp.zip
│       │   ├── zip64-2.zip
│       │   └── zip64.zip
│       ├── writer.go
│       ├── writer_test.go
│       └── zip_test.go
├── bootstrap.bash
├── bufio
│   ├── bufio.go
│   ├── bufio_test.go
│   ├── example_test.go
│   ├── export_test.go
│   ├── scan.go
│   └── scan_test.go
├── buildall.bash
├── builtin
│   └── builtin.go
├── bytes
│   ├── boundary_test.go
│   ├── buffer.go
│   ├── buffer_test.go
│   ├── bytes.go
│   ├── bytes_test.go
│   ├── compare_test.go
│   ├── example_test.go
│   ├── export_test.go
│   ├── reader.go
│   └── reader_test.go
├── clean.bash
├── clean.bat
├── clean.rc
├── cmd
│   ├── README.vendor
│   ├── addr2line
│   │   ├── addr2line_test.go
│   │   └── main.go
│   ├── api
│   │   ├── goapi.go
│   │   ├── goapi_test.go
│   │   ├── run.go
│   │   └── testdata
│   │       └── src
│   │           ├── issue21181
│   │           │   ├── dep
│   │           │   │   ├── p.go
│   │           │   │   └── p_amd64.go
│   │           │   ├── indirect
│   │           │   │   └── p.go
│   │           │   └── p
│   │           │       ├── p.go
│   │           │       ├── p_amd64.go
│   │           │       └── p_generic.go
│   │           ├── issue29837
│   │           │   └── p
│   │           │       └── README
│   │           └── pkg
│   │               ├── p1
│   │               │   ├── golden.txt
│   │               │   └── p1.go
│   │               ├── p2
│   │               │   ├── golden.txt
│   │               │   └── p2.go
│   │               └── p3
│   │                   ├── golden.txt
│   │                   └── p3.go
│   ├── asm
│   │   ├── doc.go
│   │   ├── internal
│   │   │   ├── arch
│   │   │   │   ├── arch.go
│   │   │   │   ├── arm.go
│   │   │   │   ├── arm64.go
│   │   │   │   ├── mips.go
│   │   │   │   ├── ppc64.go
│   │   │   │   └── s390x.go
│   │   │   ├── asm
│   │   │   │   ├── asm.go
│   │   │   │   ├── endtoend_test.go
│   │   │   │   ├── expr_test.go
│   │   │   │   ├── line_test.go
│   │   │   │   ├── operand_test.go
│   │   │   │   ├── parse.go
│   │   │   │   ├── pseudo_test.go
│   │   │   │   └── testdata
│   │   │   │       ├── 386.s
│   │   │   │       ├── 386enc.s
│   │   │   │       ├── amd64.s
│   │   │   │       ├── amd64enc.s
│   │   │   │       ├── amd64enc_extra.s
│   │   │   │       ├── amd64error.s
│   │   │   │       ├── arm.s
│   │   │   │       ├── arm64.s
│   │   │   │       ├── arm64enc.s
│   │   │   │       ├── arm64error.s
│   │   │   │       ├── armerror.s
│   │   │   │       ├── armv6.s
│   │   │   │       ├── avx512enc
│   │   │   │       │   ├── aes_avx512f.s
│   │   │   │       │   ├── avx512_4fmaps.s
│   │   │   │       │   ├── avx512_4vnniw.s
│   │   │   │       │   ├── avx512_bitalg.s
│   │   │   │       │   ├── avx512_ifma.s
│   │   │   │       │   ├── avx512_vbmi.s
│   │   │   │       │   ├── avx512_vbmi2.s
│   │   │   │       │   ├── avx512_vnni.s
│   │   │   │       │   ├── avx512_vpopcntdq.s
│   │   │   │       │   ├── avx512bw.s
│   │   │   │       │   ├── avx512cd.s
│   │   │   │       │   ├── avx512dq.s
│   │   │   │       │   ├── avx512er.s
│   │   │   │       │   ├── avx512f.s
│   │   │   │       │   ├── avx512pf.s
│   │   │   │       │   ├── gfni_avx512f.s
│   │   │   │       │   └── vpclmulqdq_avx512f.s
│   │   │   │       ├── mips.s
│   │   │   │       ├── mips64.s
│   │   │   │       ├── ppc64.s
│   │   │   │       ├── ppc64enc.s
│   │   │   │       └── s390x.s
│   │   │   ├── flags
│   │   │   │   └── flags.go
│   │   │   └── lex
│   │   │       ├── input.go
│   │   │       ├── lex.go
│   │   │       ├── lex_test.go
│   │   │       ├── slice.go
│   │   │       ├── stack.go
│   │   │       └── tokenizer.go
│   │   └── main.go
│   ├── buildid
│   │   ├── buildid.go
│   │   └── doc.go
│   ├── cgo
│   │   ├── ast.go
│   │   ├── doc.go
│   │   ├── gcc.go
│   │   ├── godefs.go
│   │   ├── main.go
│   │   ├── out.go
│   │   ├── util.go
│   │   └── zdefaultcc.go
│   ├── compile
│   │   ├── README.md
│   │   ├── doc.go
│   │   ├── fmt_test.go
│   │   ├── fmtmap_test.go
│   │   ├── internal
│   │   │   ├── amd64
│   │   │   │   ├── galign.go
│   │   │   │   ├── ggen.go
│   │   │   │   └── ssa.go
│   │   │   ├── arm
│   │   │   │   ├── galign.go
│   │   │   │   ├── ggen.go
│   │   │   │   └── ssa.go
│   │   │   ├── arm64
│   │   │   │   ├── galign.go
│   │   │   │   ├── ggen.go
│   │   │   │   └── ssa.go
│   │   │   ├── gc
│   │   │   │   ├── alg.go
│   │   │   │   ├── align.go
│   │   │   │   ├── bexport.go
│   │   │   │   ├── bimport.go
│   │   │   │   ├── bitset.go
│   │   │   │   ├── bootstrap.go
│   │   │   │   ├── builtin
│   │   │   │   │   └── runtime.go
│   │   │   │   ├── builtin.go
│   │   │   │   ├── builtin_test.go
│   │   │   │   ├── bv.go
│   │   │   │   ├── class_string.go
│   │   │   │   ├── closure.go
│   │   │   │   ├── const.go
│   │   │   │   ├── constFold_test.go
│   │   │   │   ├── dcl.go
│   │   │   │   ├── dep_test.go
│   │   │   │   ├── dump.go
│   │   │   │   ├── dwinl.go
│   │   │   │   ├── esc.go
│   │   │   │   ├── escape.go
│   │   │   │   ├── export.go
│   │   │   │   ├── fixedbugs_test.go
│   │   │   │   ├── float_test.go
│   │   │   │   ├── fmt.go
│   │   │   │   ├── gen.go
│   │   │   │   ├── global_test.go
│   │   │   │   ├── go.go
│   │   │   │   ├── gsubr.go
│   │   │   │   ├── iexport.go
│   │   │   │   ├── iface_test.go
│   │   │   │   ├── iimport.go
│   │   │   │   ├── init.go
│   │   │   │   ├── initorder.go
│   │   │   │   ├── inl.go
│   │   │   │   ├── inl_test.go
│   │   │   │   ├── lang_test.go
│   │   │   │   ├── lex.go
│   │   │   │   ├── lex_test.go
│   │   │   │   ├── logic_test.go
│   │   │   │   ├── main.go
│   │   │   │   ├── mapfile_mmap.go
│   │   │   │   ├── mapfile_read.go
│   │   │   │   ├── mkbuiltin.go
│   │   │   │   ├── mpfloat.go
│   │   │   │   ├── mpint.go
│   │   │   │   ├── noder.go
│   │   │   │   ├── obj.go
│   │   │   │   ├── op_string.go
│   │   │   │   ├── order.go
│   │   │   │   ├── pgen.go
│   │   │   │   ├── pgen_test.go
│   │   │   │   ├── phi.go
│   │   │   │   ├── plive.go
│   │   │   │   ├── pprof.go
│   │   │   │   ├── racewalk.go
│   │   │   │   ├── range.go
│   │   │   │   ├── reflect.go
│   │   │   │   ├── reproduciblebuilds_test.go
│   │   │   │   ├── scc.go
│   │   │   │   ├── scope.go
│   │   │   │   ├── scope_test.go
│   │   │   │   ├── select.go
│   │   │   │   ├── shift_test.go
│   │   │   │   ├── sinit.go
│   │   │   │   ├── sizeof_test.go
│   │   │   │   ├── ssa.go
│   │   │   │   ├── ssa_test.go
│   │   │   │   ├── subr.go
│   │   │   │   ├── swt.go
│   │   │   │   ├── swt_test.go
│   │   │   │   ├── syntax.go
│   │   │   │   ├── testdata
│   │   │   │   │   ├── addressed_test.go
│   │   │   │   │   ├── append_test.go
│   │   │   │   │   ├── arithBoundary_test.go
│   │   │   │   │   ├── arithConst_test.go
│   │   │   │   │   ├── arith_test.go
│   │   │   │   │   ├── array_test.go
│   │   │   │   │   ├── assert_test.go
│   │   │   │   │   ├── break_test.go
│   │   │   │   │   ├── chan_test.go
│   │   │   │   │   ├── closure_test.go
│   │   │   │   │   ├── cmpConst_test.go
│   │   │   │   │   ├── cmp_test.go
│   │   │   │   │   ├── compound_test.go
│   │   │   │   │   ├── copy_test.go
│   │   │   │   │   ├── ctl_test.go
│   │   │   │   │   ├── deferNoReturn_test.go
│   │   │   │   │   ├── divbyzero_test.go
│   │   │   │   │   ├── dupLoad_test.go
│   │   │   │   │   ├── flowgraph_generator1.go
│   │   │   │   │   ├── fp_test.go
│   │   │   │   │   ├── gen
│   │   │   │   │   │   ├── arithBoundaryGen.go
│   │   │   │   │   │   ├── arithConstGen.go
│   │   │   │   │   │   ├── cmpConstGen.go
│   │   │   │   │   │   ├── constFoldGen.go
│   │   │   │   │   │   ├── copyGen.go
│   │   │   │   │   │   └── zeroGen.go
│   │   │   │   │   ├── loadstore_test.go
│   │   │   │   │   ├── map_test.go
│   │   │   │   │   ├── namedReturn_test.go
│   │   │   │   │   ├── phi_test.go
│   │   │   │   │   ├── regalloc_test.go
│   │   │   │   │   ├── reproducible
│   │   │   │   │   │   ├── issue20272.go
│   │   │   │   │   │   ├── issue27013.go
│   │   │   │   │   │   └── issue30202.go
│   │   │   │   │   ├── short_test.go
│   │   │   │   │   ├── slice_test.go
│   │   │   │   │   ├── sqrtConst_test.go
│   │   │   │   │   ├── string_test.go
│   │   │   │   │   ├── unsafe_test.go
│   │   │   │   │   └── zero_test.go
│   │   │   │   ├── timings.go
│   │   │   │   ├── trace.go
│   │   │   │   ├── truncconst_test.go
│   │   │   │   ├── typecheck.go
│   │   │   │   ├── types.go
│   │   │   │   ├── types_acc.go
│   │   │   │   ├── universe.go
│   │   │   │   ├── unsafe.go
│   │   │   │   ├── util.go
│   │   │   │   └── walk.go
│   │   │   ├── mips
│   │   │   │   ├── galign.go
│   │   │   │   ├── ggen.go
│   │   │   │   └── ssa.go
│   │   │   ├── mips64
│   │   │   │   ├── galign.go
│   │   │   │   ├── ggen.go
│   │   │   │   └── ssa.go
│   │   │   ├── ppc64
│   │   │   │   ├── galign.go
│   │   │   │   ├── ggen.go
│   │   │   │   ├── opt.go
│   │   │   │   └── ssa.go
│   │   │   ├── s390x
│   │   │   │   ├── galign.go
│   │   │   │   ├── ggen.go
│   │   │   │   └── ssa.go
│   │   │   ├── ssa
│   │   │   │   ├── README.md
│   │   │   │   ├── TODO
│   │   │   │   ├── biasedsparsemap.go
│   │   │   │   ├── block.go
│   │   │   │   ├── branchelim.go
│   │   │   │   ├── branchelim_test.go
│   │   │   │   ├── cache.go
│   │   │   │   ├── check.go
│   │   │   │   ├── checkbce.go
│   │   │   │   ├── compile.go
│   │   │   │   ├── config.go
│   │   │   │   ├── copyelim.go
│   │   │   │   ├── copyelim_test.go
│   │   │   │   ├── critical.go
│   │   │   │   ├── cse.go
│   │   │   │   ├── cse_test.go
│   │   │   │   ├── deadcode.go
│   │   │   │   ├── deadcode_test.go
│   │   │   │   ├── deadstore.go
│   │   │   │   ├── deadstore_test.go
│   │   │   │   ├── debug.go
│   │   │   │   ├── debug_test.go
│   │   │   │   ├── decompose.go
│   │   │   │   ├── dom.go
│   │   │   │   ├── dom_test.go
│   │   │   │   ├── export_test.go
│   │   │   │   ├── flagalloc.go
│   │   │   │   ├── func.go
│   │   │   │   ├── func_test.go
│   │   │   │   ├── fuse.go
│   │   │   │   ├── fuse_test.go
│   │   │   │   ├── gen
│   │   │   │   │   ├── 386.rules
│   │   │   │   │   ├── 386Ops.go
│   │   │   │   │   ├── 386splitload.rules
│   │   │   │   │   ├── AMD64.rules
│   │   │   │   │   ├── AMD64Ops.go
│   │   │   │   │   ├── AMD64splitload.rules
│   │   │   │   │   ├── ARM.rules
│   │   │   │   │   ├── ARM64.rules
│   │   │   │   │   ├── ARM64Ops.go
│   │   │   │   │   ├── ARMOps.go
│   │   │   │   │   ├── MIPS.rules
│   │   │   │   │   ├── MIPS64.rules
│   │   │   │   │   ├── MIPS64Ops.go
│   │   │   │   │   ├── MIPSOps.go
│   │   │   │   │   ├── PPC64.rules
│   │   │   │   │   ├── PPC64Ops.go
│   │   │   │   │   ├── README
│   │   │   │   │   ├── S390X.rules
│   │   │   │   │   ├── S390XOps.go
│   │   │   │   │   ├── Wasm.rules
│   │   │   │   │   ├── WasmOps.go
│   │   │   │   │   ├── dec.rules
│   │   │   │   │   ├── dec64.rules
│   │   │   │   │   ├── dec64Ops.go
│   │   │   │   │   ├── decArgs.rules
│   │   │   │   │   ├── decArgsOps.go
│   │   │   │   │   ├── decOps.go
│   │   │   │   │   ├── generic.rules
│   │   │   │   │   ├── genericOps.go
│   │   │   │   │   ├── main.go
│   │   │   │   │   └── rulegen.go
│   │   │   │   ├── html.go
│   │   │   │   ├── id.go
│   │   │   │   ├── layout.go
│   │   │   │   ├── lca.go
│   │   │   │   ├── lca_test.go
│   │   │   │   ├── likelyadjust.go
│   │   │   │   ├── location.go
│   │   │   │   ├── loop_test.go
│   │   │   │   ├── loopbce.go
│   │   │   │   ├── loopreschedchecks.go
│   │   │   │   ├── looprotate.go
│   │   │   │   ├── lower.go
│   │   │   │   ├── magic.go
│   │   │   │   ├── magic_test.go
│   │   │   │   ├── nilcheck.go
│   │   │   │   ├── nilcheck_test.go
│   │   │   │   ├── numberlines.go
│   │   │   │   ├── op.go
│   │   │   │   ├── opGen.go
│   │   │   │   ├── opt.go
│   │   │   │   ├── passbm_test.go
│   │   │   │   ├── phielim.go
│   │   │   │   ├── phiopt.go
│   │   │   │   ├── poset.go
│   │   │   │   ├── poset_test.go
│   │   │   │   ├── print.go
│   │   │   │   ├── prove.go
│   │   │   │   ├── redblack32.go
│   │   │   │   ├── redblack32_test.go
│   │   │   │   ├── regalloc.go
│   │   │   │   ├── regalloc_test.go
│   │   │   │   ├── rewrite.go
│   │   │   │   ├── rewrite386.go
│   │   │   │   ├── rewrite386splitload.go
│   │   │   │   ├── rewriteAMD64.go
│   │   │   │   ├── rewriteAMD64splitload.go
│   │   │   │   ├── rewriteARM.go
│   │   │   │   ├── rewriteARM64.go
│   │   │   │   ├── rewriteMIPS.go
│   │   │   │   ├── rewriteMIPS64.go
│   │   │   │   ├── rewritePPC64.go
│   │   │   │   ├── rewriteS390X.go
│   │   │   │   ├── rewriteWasm.go
│   │   │   │   ├── rewrite_test.go
│   │   │   │   ├── rewritedec.go
│   │   │   │   ├── rewritedec64.go
│   │   │   │   ├── rewritedecArgs.go
│   │   │   │   ├── rewritegeneric.go
│   │   │   │   ├── schedule.go
│   │   │   │   ├── schedule_test.go
│   │   │   │   ├── shift_test.go
│   │   │   │   ├── shortcircuit.go
│   │   │   │   ├── shortcircuit_test.go
│   │   │   │   ├── sizeof_test.go
│   │   │   │   ├── softfloat.go
│   │   │   │   ├── sparsemap.go
│   │   │   │   ├── sparseset.go
│   │   │   │   ├── sparsetree.go
│   │   │   │   ├── sparsetreemap.go
│   │   │   │   ├── stackalloc.go
│   │   │   │   ├── stackframe.go
│   │   │   │   ├── stmtlines_test.go
│   │   │   │   ├── testdata
│   │   │   │   │   ├── hist.dlv-dbg.nexts
│   │   │   │   │   ├── hist.dlv-opt.nexts
│   │   │   │   │   ├── hist.gdb-dbg.nexts
│   │   │   │   │   ├── hist.gdb-opt.nexts
│   │   │   │   │   ├── hist.go
│   │   │   │   │   ├── i22558.dlv-dbg.nexts
│   │   │   │   │   ├── i22558.gdb-dbg.nexts
│   │   │   │   │   ├── i22558.go
│   │   │   │   │   ├── i22600.dlv-dbg-race.nexts
│   │   │   │   │   ├── i22600.gdb-dbg-race.nexts
│   │   │   │   │   ├── i22600.go
│   │   │   │   │   ├── infloop.dlv-opt.nexts
│   │   │   │   │   ├── infloop.gdb-opt.nexts
│   │   │   │   │   ├── infloop.go
│   │   │   │   │   ├── scopes.dlv-dbg.nexts
│   │   │   │   │   ├── scopes.dlv-opt.nexts
│   │   │   │   │   ├── scopes.gdb-dbg.nexts
│   │   │   │   │   ├── scopes.gdb-opt.nexts
│   │   │   │   │   └── scopes.go
│   │   │   │   ├── tighten.go
│   │   │   │   ├── trim.go
│   │   │   │   ├── value.go
│   │   │   │   ├── writebarrier.go
│   │   │   │   ├── writebarrier_test.go
│   │   │   │   ├── xposmap.go
│   │   │   │   ├── zcse.go
│   │   │   │   └── zeroextension_test.go
│   │   │   ├── syntax
│   │   │   │   ├── branches.go
│   │   │   │   ├── dumper.go
│   │   │   │   ├── dumper_test.go
│   │   │   │   ├── error_test.go
│   │   │   │   ├── nodes.go
│   │   │   │   ├── nodes_test.go
│   │   │   │   ├── operator_string.go
│   │   │   │   ├── parser.go
│   │   │   │   ├── parser_test.go
│   │   │   │   ├── pos.go
│   │   │   │   ├── printer.go
│   │   │   │   ├── printer_test.go
│   │   │   │   ├── scanner.go
│   │   │   │   ├── scanner_test.go
│   │   │   │   ├── source.go
│   │   │   │   ├── syntax.go
│   │   │   │   ├── testdata
│   │   │   │   │   ├── issue20789.src
│   │   │   │   │   ├── issue23385.src
│   │   │   │   │   ├── issue23434.src
│   │   │   │   │   ├── issue31092.src
│   │   │   │   │   └── sample.src
│   │   │   │   ├── token_string.go
│   │   │   │   └── tokens.go
│   │   │   ├── test
│   │   │   │   ├── README
│   │   │   │   ├── divconst_test.go
│   │   │   │   └── test.go
│   │   │   ├── types
│   │   │   │   ├── etype_string.go
│   │   │   │   ├── identity.go
│   │   │   │   ├── pkg.go
│   │   │   │   ├── scope.go
│   │   │   │   ├── sizeof_test.go
│   │   │   │   ├── sym.go
│   │   │   │   ├── sym_test.go
│   │   │   │   ├── type.go
│   │   │   │   └── utils.go
│   │   │   ├── wasm
│   │   │   │   └── ssa.go
│   │   │   └── x86
│   │   │       ├── 387.go
│   │   │       ├── galign.go
│   │   │       ├── ggen.go
│   │   │       └── ssa.go
│   │   └── main.go
│   ├── cover
│   │   ├── cover.go
│   │   ├── cover_test.go
│   │   ├── doc.go
│   │   ├── func.go
│   │   ├── html.go
│   │   ├── profile.go
│   │   └── testdata
│   │       ├── directives.go
│   │       ├── html
│   │       │   ├── html.go
│   │       │   ├── html.golden
│   │       │   └── html_test.go
│   │       ├── main.go
│   │       ├── p.go
│   │       ├── profile.cov
│   │       ├── test.go
│   │       └── toolexec.go
│   ├── dist
│   │   ├── README
│   │   ├── build.go
│   │   ├── buildgo.go
│   │   ├── buildruntime.go
│   │   ├── buildtool.go
│   │   ├── cpuid_386.s
│   │   ├── cpuid_amd64.s
│   │   ├── cpuid_default.s
│   │   ├── doc.go
│   │   ├── imports.go
│   │   ├── main.go
│   │   ├── sys_default.go
│   │   ├── sys_windows.go
│   │   ├── test.go
│   │   ├── test_linux.go
│   │   ├── util.go
│   │   ├── util_gc.go
│   │   ├── util_gccgo.go
│   │   ├── vfp_arm.s
│   │   └── vfp_default.s
│   ├── doc
│   │   ├── dirs.go
│   │   ├── doc_test.go
│   │   ├── main.go
│   │   ├── pkg.go
│   │   └── testdata
│   │       ├── nested
│   │       │   ├── ignore.go
│   │       │   └── nested
│   │       │       └── real.go
│   │       └── pkg.go
│   ├── fix
│   │   ├── cftype.go
│   │   ├── cftype_test.go
│   │   ├── context.go
│   │   ├── context_test.go
│   │   ├── doc.go
│   │   ├── egltype.go
│   │   ├── egltype_test.go
│   │   ├── fix.go
│   │   ├── gotypes.go
│   │   ├── gotypes_test.go
│   │   ├── import_test.go
│   │   ├── jnitype.go
│   │   ├── jnitype_test.go
│   │   ├── main.go
│   │   ├── main_test.go
│   │   ├── netipv6zone.go
│   │   ├── netipv6zone_test.go
│   │   ├── printerconfig.go
│   │   ├── printerconfig_test.go
│   │   └── typecheck.go
│   ├── go
│   │   ├── alldocs.go
│   │   ├── go11.go
│   │   ├── go_test.go
│   │   ├── go_unix_test.go
│   │   ├── go_windows_test.go
│   │   ├── help_test.go
│   │   ├── init_test.go
│   │   ├── internal
│   │   │   ├── auth
│   │   │   │   ├── auth.go
│   │   │   │   ├── netrc.go
│   │   │   │   └── netrc_test.go
│   │   │   ├── base
│   │   │   │   ├── base.go
│   │   │   │   ├── env.go
│   │   │   │   ├── flag.go
│   │   │   │   ├── goflags.go
│   │   │   │   ├── path.go
│   │   │   │   ├── signal.go
│   │   │   │   ├── signal_notunix.go
│   │   │   │   ├── signal_unix.go
│   │   │   │   └── tool.go
│   │   │   ├── bug
│   │   │   │   └── bug.go
│   │   │   ├── cache
│   │   │   │   ├── cache.go
│   │   │   │   ├── cache_test.go
│   │   │   │   ├── default.go
│   │   │   │   ├── hash.go
│   │   │   │   └── hash_test.go
│   │   │   ├── cfg
│   │   │   │   ├── cfg.go
│   │   │   │   ├── zdefaultcc.go
│   │   │   │   └── zosarch.go
│   │   │   ├── clean
│   │   │   │   └── clean.go
│   │   │   ├── cmdflag
│   │   │   │   └── flag.go
│   │   │   ├── dirhash
│   │   │   │   ├── hash.go
│   │   │   │   └── hash_test.go
│   │   │   ├── doc
│   │   │   │   └── doc.go
│   │   │   ├── envcmd
│   │   │   │   └── env.go
│   │   │   ├── fix
│   │   │   │   └── fix.go
│   │   │   ├── fmtcmd
│   │   │   │   └── fmt.go
│   │   │   ├── generate
│   │   │   │   ├── generate.go
│   │   │   │   └── generate_test.go
│   │   │   ├── get
│   │   │   │   ├── discovery.go
│   │   │   │   ├── get.go
│   │   │   │   ├── path.go
│   │   │   │   ├── pkg_test.go
│   │   │   │   ├── tag_test.go
│   │   │   │   ├── vcs.go
│   │   │   │   └── vcs_test.go
│   │   │   ├── help
│   │   │   │   ├── help.go
│   │   │   │   └── helpdoc.go
│   │   │   ├── imports
│   │   │   │   ├── build.go
│   │   │   │   ├── read.go
│   │   │   │   ├── read_test.go
│   │   │   │   ├── scan.go
│   │   │   │   ├── scan_test.go
│   │   │   │   ├── tags.go
│   │   │   │   └── testdata
│   │   │   │       ├── android
│   │   │   │       │   ├── a_android.go
│   │   │   │       │   ├── b_android_arm64.go
│   │   │   │       │   ├── c_linux.go
│   │   │   │       │   ├── d_linux_arm64.go
│   │   │   │       │   ├── e.go
│   │   │   │       │   ├── f.go
│   │   │   │       │   ├── g.go
│   │   │   │       │   ├── tags.txt
│   │   │   │       │   └── want.txt
│   │   │   │       ├── illumos
│   │   │   │       │   ├── a_illumos.go
│   │   │   │       │   ├── b_illumos_amd64.go
│   │   │   │       │   ├── c_solaris.go
│   │   │   │       │   ├── d_solaris_amd64.go
│   │   │   │       │   ├── e.go
│   │   │   │       │   ├── f.go
│   │   │   │       │   ├── g.go
│   │   │   │       │   ├── tags.txt
│   │   │   │       │   └── want.txt
│   │   │   │       └── star
│   │   │   │           ├── tags.txt
│   │   │   │           ├── want.txt
│   │   │   │           ├── x.go
│   │   │   │           ├── x1.go
│   │   │   │           ├── x_darwin.go
│   │   │   │           └── x_windows.go
│   │   │   ├── list
│   │   │   │   ├── context.go
│   │   │   │   └── list.go
│   │   │   ├── load
│   │   │   │   ├── flag.go
│   │   │   │   ├── flag_test.go
│   │   │   │   ├── path.go
│   │   │   │   ├── pkg.go
│   │   │   │   ├── pkg_test.go
│   │   │   │   ├── search.go
│   │   │   │   └── test.go
│   │   │   ├── lockedfile
│   │   │   │   ├── internal
│   │   │   │   │   └── filelock
│   │   │   │   │       ├── filelock.go
│   │   │   │   │       ├── filelock_fcntl.go
│   │   │   │   │       ├── filelock_other.go
│   │   │   │   │       ├── filelock_plan9.go
│   │   │   │   │       ├── filelock_test.go
│   │   │   │   │       ├── filelock_unix.go
│   │   │   │   │       └── filelock_windows.go
│   │   │   │   ├── lockedfile.go
│   │   │   │   ├── lockedfile_filelock.go
│   │   │   │   ├── lockedfile_plan9.go
│   │   │   │   ├── lockedfile_test.go
│   │   │   │   └── mutex.go
│   │   │   ├── modcmd
│   │   │   │   ├── download.go
│   │   │   │   ├── edit.go
│   │   │   │   ├── graph.go
│   │   │   │   ├── init.go
│   │   │   │   ├── mod.go
│   │   │   │   ├── tidy.go
│   │   │   │   ├── vendor.go
│   │   │   │   ├── verify.go
│   │   │   │   └── why.go
│   │   │   ├── modconv
│   │   │   │   ├── convert.go
│   │   │   │   ├── convert_test.go
│   │   │   │   ├── dep.go
│   │   │   │   ├── glide.go
│   │   │   │   ├── glock.go
│   │   │   │   ├── godeps.go
│   │   │   │   ├── modconv.go
│   │   │   │   ├── modconv_test.go
│   │   │   │   ├── testdata
│   │   │   │   │   ├── cockroach.glock
│   │   │   │   │   ├── cockroach.out
│   │   │   │   │   ├── dockermachine.godeps
│   │   │   │   │   ├── dockermachine.out
│   │   │   │   │   ├── dockerman.glide
│   │   │   │   │   ├── dockerman.out
│   │   │   │   │   ├── govmomi.out
│   │   │   │   │   ├── govmomi.vmanifest
│   │   │   │   │   ├── juju.out
│   │   │   │   │   ├── juju.tsv
│   │   │   │   │   ├── moby.out
│   │   │   │   │   ├── moby.vconf
│   │   │   │   │   ├── panicparse.out
│   │   │   │   │   ├── panicparse.vyml
│   │   │   │   │   ├── prometheus.out
│   │   │   │   │   ├── prometheus.vjson
│   │   │   │   │   ├── traefik.dep
│   │   │   │   │   ├── traefik.out
│   │   │   │   │   ├── upspin.dep
│   │   │   │   │   └── upspin.out
│   │   │   │   ├── tsv.go
│   │   │   │   ├── vconf.go
│   │   │   │   ├── vjson.go
│   │   │   │   ├── vmanifest.go
│   │   │   │   └── vyml.go
│   │   │   ├── modfetch
│   │   │   │   ├── bootstrap.go
│   │   │   │   ├── cache.go
│   │   │   │   ├── cache_test.go
│   │   │   │   ├── codehost
│   │   │   │   │   ├── codehost.go
│   │   │   │   │   ├── git.go
│   │   │   │   │   ├── git_test.go
│   │   │   │   │   ├── shell.go
│   │   │   │   │   └── vcs.go
│   │   │   │   ├── coderepo.go
│   │   │   │   ├── coderepo_test.go
│   │   │   │   ├── fetch.go
│   │   │   │   ├── key.go
│   │   │   │   ├── proxy.go
│   │   │   │   ├── pseudo.go
│   │   │   │   ├── pseudo_test.go
│   │   │   │   ├── repo.go
│   │   │   │   ├── sumdb.go
│   │   │   │   └── unzip.go
│   │   │   ├── modfile
│   │   │   │   ├── gopkgin.go
│   │   │   │   ├── print.go
│   │   │   │   ├── read.go
│   │   │   │   ├── read_test.go
│   │   │   │   ├── rule.go
│   │   │   │   ├── rule_test.go
│   │   │   │   └── testdata
│   │   │   │       ├── block.golden
│   │   │   │       ├── block.in
│   │   │   │       ├── comment.golden
│   │   │   │       ├── comment.in
│   │   │   │       ├── empty.golden
│   │   │   │       ├── empty.in
│   │   │   │       ├── gopkg.in.golden
│   │   │   │       ├── module.golden
│   │   │   │       ├── module.in
│   │   │   │       ├── replace.golden
│   │   │   │       ├── replace.in
│   │   │   │       ├── replace2.golden
│   │   │   │       ├── replace2.in
│   │   │   │       └── rule1.golden
│   │   │   ├── modget
│   │   │   │   └── get.go
│   │   │   ├── modinfo
│   │   │   │   └── info.go
│   │   │   ├── modload
│   │   │   │   ├── build.go
│   │   │   │   ├── help.go
│   │   │   │   ├── import.go
│   │   │   │   ├── import_test.go
│   │   │   │   ├── init.go
│   │   │   │   ├── list.go
│   │   │   │   ├── load.go
│   │   │   │   ├── query.go
│   │   │   │   ├── query_test.go
│   │   │   │   ├── search.go
│   │   │   │   └── testgo.go
│   │   │   ├── module
│   │   │   │   ├── module.go
│   │   │   │   └── module_test.go
│   │   │   ├── mvs
│   │   │   │   ├── mvs.go
│   │   │   │   └── mvs_test.go
│   │   │   ├── note
│   │   │   │   ├── example_test.go
│   │   │   │   ├── note.go
│   │   │   │   └── note_test.go
│   │   │   ├── par
│   │   │   │   ├── work.go
│   │   │   │   └── work_test.go
│   │   │   ├── renameio
│   │   │   │   ├── renameio.go
│   │   │   │   ├── renameio_test.go
│   │   │   │   └── umask_test.go
│   │   │   ├── robustio
│   │   │   │   ├── robustio.go
│   │   │   │   ├── robustio_other.go
│   │   │   │   └── robustio_windows.go
│   │   │   ├── run
│   │   │   │   └── run.go
│   │   │   ├── search
│   │   │   │   ├── search.go
│   │   │   │   └── search_test.go
│   │   │   ├── semver
│   │   │   │   ├── semver.go
│   │   │   │   └── semver_test.go
│   │   │   ├── str
│   │   │   │   ├── path.go
│   │   │   │   └── str.go
│   │   │   ├── sumweb
│   │   │   │   ├── cache.go
│   │   │   │   ├── client.go
│   │   │   │   ├── client_test.go
│   │   │   │   ├── encode.go
│   │   │   │   ├── encode_test.go
│   │   │   │   ├── server.go
│   │   │   │   └── test.go
│   │   │   ├── test
│   │   │   │   ├── cover.go
│   │   │   │   ├── test.go
│   │   │   │   └── testflag.go
│   │   │   ├── tlog
│   │   │   │   ├── ct_test.go
│   │   │   │   ├── note.go
│   │   │   │   ├── note_test.go
│   │   │   │   ├── tile.go
│   │   │   │   ├── tlog.go
│   │   │   │   └── tlog_test.go
│   │   │   ├── tool
│   │   │   │   └── tool.go
│   │   │   ├── txtar
│   │   │   │   ├── archive.go
│   │   │   │   └── archive_test.go
│   │   │   ├── version
│   │   │   │   ├── exe.go
│   │   │   │   └── version.go
│   │   │   ├── vet
│   │   │   │   ├── vet.go
│   │   │   │   └── vetflag.go
│   │   │   ├── web
│   │   │   │   ├── api.go
│   │   │   │   ├── bootstrap.go
│   │   │   │   ├── file_test.go
│   │   │   │   ├── http.go
│   │   │   │   ├── url.go
│   │   │   │   ├── url_other.go
│   │   │   │   ├── url_other_test.go
│   │   │   │   ├── url_test.go
│   │   │   │   ├── url_windows.go
│   │   │   │   └── url_windows_test.go
│   │   │   └── work
│   │   │       ├── action.go
│   │   │       ├── build.go
│   │   │       ├── build_test.go
│   │   │       ├── buildid.go
│   │   │       ├── exec.go
│   │   │       ├── gc.go
│   │   │       ├── gccgo.go
│   │   │       ├── init.go
│   │   │       ├── security.go
│   │   │       ├── security_test.go
│   │   │       └── testgo.go
│   │   ├── main.go
│   │   ├── mkalldocs.sh
│   │   ├── note_test.go
│   │   ├── proxy_test.go
│   │   ├── script_test.go
│   │   ├── testdata
│   │   │   ├── addmod.go
│   │   │   ├── badmod
│   │   │   │   ├── go.mod
│   │   │   │   └── x.go
│   │   │   ├── dep_test.go
│   │   │   ├── example1_test.go
│   │   │   ├── example2_test.go
│   │   │   ├── failssh
│   │   │   │   └── ssh
│   │   │   ├── flag_test.go
│   │   │   ├── generate
│   │   │   │   ├── test1.go
│   │   │   │   ├── test2.go
│   │   │   │   ├── test3.go
│   │   │   │   └── test4.go
│   │   │   ├── importcom
│   │   │   │   ├── bad.go
│   │   │   │   ├── conflict.go
│   │   │   │   ├── src
│   │   │   │   │   ├── bad
│   │   │   │   │   │   └── bad.go
│   │   │   │   │   ├── conflict
│   │   │   │   │   │   ├── a.go
│   │   │   │   │   │   └── b.go
│   │   │   │   │   ├── works
│   │   │   │   │   │   └── x
│   │   │   │   │   │       ├── x.go
│   │   │   │   │   │       └── x1.go
│   │   │   │   │   └── wrongplace
│   │   │   │   │       └── x.go
│   │   │   │   ├── works.go
│   │   │   │   └── wrongplace.go
│   │   │   ├── importcycle
│   │   │   │   └── src
│   │   │   │       └── selfimport
│   │   │   │           └── selfimport.go
│   │   │   ├── local
│   │   │   │   ├── easy.go
│   │   │   │   ├── easysub
│   │   │   │   │   ├── easysub.go
│   │   │   │   │   └── main.go
│   │   │   │   ├── hard.go
│   │   │   │   └── sub
│   │   │   │       ├── sub
│   │   │   │       │   └── subsub.go
│   │   │   │       └── sub.go
│   │   │   ├── mod
│   │   │   │   ├── README
│   │   │   │   ├── example.com_badchain_a_v1.0.0.txt
│   │   │   │   ├── example.com_badchain_a_v1.1.0.txt
│   │   │   │   ├── example.com_badchain_b_v1.0.0.txt
│   │   │   │   ├── example.com_badchain_b_v1.1.0.txt
│   │   │   │   ├── example.com_badchain_c_v1.0.0.txt
│   │   │   │   ├── example.com_badchain_c_v1.1.0.txt
│   │   │   │   ├── example.com_dotgo.go_v1.0.0.txt
│   │   │   │   ├── example.com_downgrade_v2.0.0.txt
│   │   │   │   ├── example.com_downgrade_v2_v2.0.1.txt
│   │   │   │   ├── example.com_invalidpath_v1_v1.0.0.txt
│   │   │   │   ├── example.com_join_subpkg_v1.0.0.txt
│   │   │   │   ├── example.com_join_subpkg_v1.1.0.txt
│   │   │   │   ├── example.com_join_v1.0.0.txt
│   │   │   │   ├── example.com_join_v1.1.0.txt
│   │   │   │   ├── example.com_latemigrate_v2_v2.0.0.txt
│   │   │   │   ├── example.com_latemigrate_v2_v2.0.1.txt
│   │   │   │   ├── example.com_missingpkg_v1.0.0.txt
│   │   │   │   ├── example.com_missingpkg_v1.0.1-beta.txt
│   │   │   │   ├── example.com_nest_sub_v1.0.0.txt
│   │   │   │   ├── example.com_nest_v1.0.0.txt
│   │   │   │   ├── example.com_nest_v1.1.0.txt
│   │   │   │   ├── example.com_newcycle_a_v1.0.0.txt
│   │   │   │   ├── example.com_newcycle_a_v1.0.1.txt
│   │   │   │   ├── example.com_newcycle_b_v1.0.0.txt
│   │   │   │   ├── example.com_noroot_v1.0.0.txt
│   │   │   │   ├── example.com_noroot_v1.0.1.txt
│   │   │   │   ├── example.com_notags_v0.0.0-20190507143103-cc8cbe209b64.txt
│   │   │   │   ├── example.com_printversion_v0.1.0.txt
│   │   │   │   ├── example.com_printversion_v1.0.0.txt
│   │   │   │   ├── example.com_pseudoupgrade_v0.0.0-20190430073000-30950c05d534.txt
│   │   │   │   ├── example.com_pseudoupgrade_v0.1.0.txt
│   │   │   │   ├── example.com_pseudoupgrade_v0.1.1-0.20190429073117-b5426c86b553.txt
│   │   │   │   ├── example.com_split_subpkg_v1.1.0.txt
│   │   │   │   ├── example.com_split_v1.0.0.txt
│   │   │   │   ├── example.com_split_v1.1.0.txt
│   │   │   │   ├── example.com_tools_v1.0.0.txt
│   │   │   │   ├── example.com_usemissingpre_v1.0.0.txt
│   │   │   │   ├── example.com_v1.0.0.txt
│   │   │   │   ├── example.com_version_v1.0.0.txt
│   │   │   │   ├── example.com_version_v1.0.1.txt
│   │   │   │   ├── example.com_version_v1.1.0.txt
│   │   │   │   ├── github.com_dmitshur-test_modtest5_v0.0.0-20190619020302-197a620e0c9a.txt
│   │   │   │   ├── github.com_dmitshur-test_modtest5_v0.5.0-alpha.0.20190619023908-3da23a9deb9e.txt
│   │   │   │   ├── github.com_dmitshur-test_modtest5_v0.5.0-alpha.txt
│   │   │   │   ├── golang.org_notx_useinternal_v0.1.0.txt
│   │   │   │   ├── golang.org_x_internal_v0.1.0.txt
│   │   │   │   ├── golang.org_x_text_v0.0.0-20170915032832-14c0d48ead0c.txt
│   │   │   │   ├── golang.org_x_text_v0.3.0.txt
│   │   │   │   ├── golang.org_x_useinternal_v0.1.0.txt
│   │   │   │   ├── gopkg.in_dummy.v2-unstable_v2.0.0.txt
│   │   │   │   ├── patch.example.com_depofdirectpatch_v1.0.0.txt
│   │   │   │   ├── patch.example.com_depofdirectpatch_v1.0.1.txt
│   │   │   │   ├── patch.example.com_direct_v1.0.0.txt
│   │   │   │   ├── patch.example.com_direct_v1.0.1.txt
│   │   │   │   ├── patch.example.com_direct_v1.1.0.txt
│   │   │   │   ├── patch.example.com_indirect_v1.0.0.txt
│   │   │   │   ├── patch.example.com_indirect_v1.0.1.txt
│   │   │   │   ├── patch.example.com_indirect_v1.1.0.txt
│   │   │   │   ├── rsc.io_!c!g!o_v1.0.0.txt
│   │   │   │   ├── rsc.io_!q!u!o!t!e_v1.5.2.txt
│   │   │   │   ├── rsc.io_!q!u!o!t!e_v1.5.3-!p!r!e.txt
│   │   │   │   ├── rsc.io_badfile1_v1.0.0.txt
│   │   │   │   ├── rsc.io_badfile2_v1.0.0.txt
│   │   │   │   ├── rsc.io_badfile3_v1.0.0.txt
│   │   │   │   ├── rsc.io_badfile4_v1.0.0.txt
│   │   │   │   ├── rsc.io_badfile5_v1.0.0.txt
│   │   │   │   ├── rsc.io_badmod_v1.0.0.txt
│   │   │   │   ├── rsc.io_badsum_v1.0.0.txt
│   │   │   │   ├── rsc.io_badsum_v1.0.1.txt
│   │   │   │   ├── rsc.io_badzip_v1.0.0.txt
│   │   │   │   ├── rsc.io_breaker_v1.0.0.txt
│   │   │   │   ├── rsc.io_breaker_v2.0.0+incompatible.txt
│   │   │   │   ├── rsc.io_breaker_v2.0.0.txt
│   │   │   │   ├── rsc.io_fortune_v1.0.0.txt
│   │   │   │   ├── rsc.io_fortune_v2_v2.0.0.txt
│   │   │   │   ├── rsc.io_quote_v0.0.0-20180214005133-e7a685a342c0.txt
│   │   │   │   ├── rsc.io_quote_v0.0.0-20180214005840-23179ee8a569.txt
│   │   │   │   ├── rsc.io_quote_v0.0.0-20180628003336-dd9747d19b04.txt
│   │   │   │   ├── rsc.io_quote_v0.0.0-20180709153244-fd906ed3b100.txt
│   │   │   │   ├── rsc.io_quote_v0.0.0-20180709160352-0d003b9c4bfa.txt
│   │   │   │   ├── rsc.io_quote_v0.0.0-20180709162749-b44a0b17b2d1.txt
│   │   │   │   ├── rsc.io_quote_v0.0.0-20180709162816-fe488b867524.txt
│   │   │   │   ├── rsc.io_quote_v0.0.0-20180709162918-a91498bed0a7.txt
│   │   │   │   ├── rsc.io_quote_v0.0.0-20180710144737-5d9f230bcfba.txt
│   │   │   │   ├── rsc.io_quote_v1.0.0.txt
│   │   │   │   ├── rsc.io_quote_v1.1.0.txt
│   │   │   │   ├── rsc.io_quote_v1.2.0.txt
│   │   │   │   ├── rsc.io_quote_v1.2.1.txt
│   │   │   │   ├── rsc.io_quote_v1.3.0.txt
│   │   │   │   ├── rsc.io_quote_v1.4.0.txt
│   │   │   │   ├── rsc.io_quote_v1.5.0.txt
│   │   │   │   ├── rsc.io_quote_v1.5.1.txt
│   │   │   │   ├── rsc.io_quote_v1.5.2.txt
│   │   │   │   ├── rsc.io_quote_v1.5.3-pre1.txt
│   │   │   │   ├── rsc.io_quote_v2.0.0.txt
│   │   │   │   ├── rsc.io_quote_v2_v2.0.1.txt
│   │   │   │   ├── rsc.io_quote_v3_v3.0.0.txt
│   │   │   │   ├── rsc.io_sampler_v1.0.0.txt
│   │   │   │   ├── rsc.io_sampler_v1.2.0.txt
│   │   │   │   ├── rsc.io_sampler_v1.2.1.txt
│   │   │   │   ├── rsc.io_sampler_v1.3.0.txt
│   │   │   │   ├── rsc.io_sampler_v1.3.1.txt
│   │   │   │   ├── rsc.io_sampler_v1.99.99.txt
│   │   │   │   └── rsc.io_testonly_v1.0.0.txt
│   │   │   ├── modlegacy
│   │   │   │   └── src
│   │   │   │       ├── new
│   │   │   │       │   ├── go.mod
│   │   │   │       │   ├── new.go
│   │   │   │       │   ├── p1
│   │   │   │       │   │   └── p1.go
│   │   │   │       │   ├── p2
│   │   │   │       │   │   └── p2.go
│   │   │   │       │   └── sub
│   │   │   │       │       ├── go.mod
│   │   │   │       │       ├── inner
│   │   │   │       │       │   ├── go.mod
│   │   │   │       │       │   └── x
│   │   │   │       │       │       └── x.go
│   │   │   │       │       └── x
│   │   │   │       │           └── v1
│   │   │   │       │               └── y
│   │   │   │       │                   └── y.go
│   │   │   │       └── old
│   │   │   │           ├── p1
│   │   │   │           │   └── p1.go
│   │   │   │           └── p2
│   │   │   │               └── p2.go
│   │   │   ├── norunexample
│   │   │   │   ├── example_test.go
│   │   │   │   └── test_test.go
│   │   │   ├── print_goroot.go
│   │   │   ├── rundir
│   │   │   │   ├── sub
│   │   │   │   │   └── sub.go
│   │   │   │   └── x.go
│   │   │   ├── savedir.go
│   │   │   ├── script
│   │   │   │   ├── README
│   │   │   │   ├── bug.txt
│   │   │   │   ├── build_GOTMPDIR.txt
│   │   │   │   ├── build_acl_windows.txt
│   │   │   │   ├── build_cache_compile.txt
│   │   │   │   ├── build_cache_gomips.txt
│   │   │   │   ├── build_cache_link.txt
│   │   │   │   ├── build_cache_output.txt
│   │   │   │   ├── build_cache_trimpath.txt
│   │   │   │   ├── build_multi_main.txt
│   │   │   │   ├── build_nocache.txt
│   │   │   │   ├── build_relative_pkgdir.txt
│   │   │   │   ├── build_relative_tmpdir.txt
│   │   │   │   ├── build_runtime_gcflags.txt
│   │   │   │   ├── build_trimpath.txt
│   │   │   │   ├── cache_unix.txt
│   │   │   │   ├── cgo_syso_issue29253.txt
│   │   │   │   ├── clean_testcache.txt
│   │   │   │   ├── cmd_import_error.txt
│   │   │   │   ├── cover_atomic_pkgall.txt
│   │   │   │   ├── cover_mod_empty.txt
│   │   │   │   ├── cover_pkgall_multiple_mains.txt
│   │   │   │   ├── cover_pkgall_runtime.txt
│   │   │   │   ├── cpu_profile_twice.txt
│   │   │   │   ├── env_write.txt
│   │   │   │   ├── fileline.txt
│   │   │   │   ├── gcflags_patterns.txt
│   │   │   │   ├── get_404_meta.txt
│   │   │   │   ├── get_brace.txt
│   │   │   │   ├── get_dotfiles.txt
│   │   │   │   ├── get_insecure_redirect.txt
│   │   │   │   ├── get_tilde.txt
│   │   │   │   ├── get_unicode.txt
│   │   │   │   ├── get_with_git_trace.txt
│   │   │   │   ├── goflags.txt
│   │   │   │   ├── gopath_std_vendor.txt
│   │   │   │   ├── help.txt
│   │   │   │   ├── install_cleans_build.txt
│   │   │   │   ├── install_cmd_gobin.txt
│   │   │   │   ├── install_cross_gobin.txt
│   │   │   │   ├── install_rebuild_gopath.txt
│   │   │   │   ├── install_rebuild_removed.txt
│   │   │   │   ├── linkname.txt
│   │   │   │   ├── list_ambiguous_path.txt
│   │   │   │   ├── list_bad_import.txt
│   │   │   │   ├── list_compiled_imports.txt
│   │   │   │   ├── list_find.txt
│   │   │   │   ├── list_importmap.txt
│   │   │   │   ├── list_split_main.txt
│   │   │   │   ├── list_std.txt
│   │   │   │   ├── list_tags.txt
│   │   │   │   ├── list_test_e.txt
│   │   │   │   ├── list_test_err.txt
│   │   │   │   ├── list_test_imports.txt
│   │   │   │   ├── list_test_non_go_files.txt
│   │   │   │   ├── mod_alt_goroot.txt
│   │   │   │   ├── mod_auth.txt
│   │   │   │   ├── mod_bad_domain.txt
│   │   │   │   ├── mod_bad_filenames.txt
│   │   │   │   ├── mod_build_tags.txt
│   │   │   │   ├── mod_build_versioned.txt
│   │   │   │   ├── mod_case.txt
│   │   │   │   ├── mod_case_cgo.txt
│   │   │   │   ├── mod_clean_cache.txt
│   │   │   │   ├── mod_concurrent.txt
│   │   │   │   ├── mod_convert_dep.txt
│   │   │   │   ├── mod_convert_git.txt
│   │   │   │   ├── mod_convert_glide.txt
│   │   │   │   ├── mod_convert_glockfile.txt
│   │   │   │   ├── mod_convert_godeps.txt
│   │   │   │   ├── mod_convert_tsv.txt
│   │   │   │   ├── mod_convert_vendor_conf.txt
│   │   │   │   ├── mod_convert_vendor_json.txt
│   │   │   │   ├── mod_convert_vendor_manifest.txt
│   │   │   │   ├── mod_convert_vendor_yml.txt
│   │   │   │   ├── mod_dir.txt
│   │   │   │   ├── mod_doc.txt
│   │   │   │   ├── mod_domain_root.txt
│   │   │   │   ├── mod_dot.txt
│   │   │   │   ├── mod_download.txt
│   │   │   │   ├── mod_download_hash.txt
│   │   │   │   ├── mod_edit.txt
│   │   │   │   ├── mod_edit_go.txt
│   │   │   │   ├── mod_enabled.txt
│   │   │   │   ├── mod_file_proxy.txt
│   │   │   │   ├── mod_find.txt
│   │   │   │   ├── mod_fs_patterns.txt
│   │   │   │   ├── mod_get_cmd.txt
│   │   │   │   ├── mod_get_commit.txt
│   │   │   │   ├── mod_get_direct.txt
│   │   │   │   ├── mod_get_downgrade.txt
│   │   │   │   ├── mod_get_fallback.txt
│   │   │   │   ├── mod_get_hash.txt
│   │   │   │   ├── mod_get_incompatible.txt
│   │   │   │   ├── mod_get_indirect.txt
│   │   │   │   ├── mod_get_insecure_redirect.txt
│   │   │   │   ├── mod_get_latest_pseudo.txt
│   │   │   │   ├── mod_get_local.txt
│   │   │   │   ├── mod_get_main.txt
│   │   │   │   ├── mod_get_major.txt
│   │   │   │   ├── mod_get_moved.txt
│   │   │   │   ├── mod_get_newcycle.txt
│   │   │   │   ├── mod_get_none.txt
│   │   │   │   ├── mod_get_patterns.txt
│   │   │   │   ├── mod_get_private_vcs.txt
│   │   │   │   ├── mod_get_pseudo.txt
│   │   │   │   ├── mod_get_pseudo_other_branch.txt
│   │   │   │   ├── mod_get_pseudo_prefix.txt
│   │   │   │   ├── mod_get_svn.txt
│   │   │   │   ├── mod_get_tags.txt
│   │   │   │   ├── mod_get_test.txt
│   │   │   │   ├── mod_get_trailing_slash.txt
│   │   │   │   ├── mod_get_upgrade.txt
│   │   │   │   ├── mod_get_upgrade_pseudo.txt
│   │   │   │   ├── mod_getmode_vendor.txt
│   │   │   │   ├── mod_git_export_subst.txt
│   │   │   │   ├── mod_go_version.txt
│   │   │   │   ├── mod_gobuild_import.txt
│   │   │   │   ├── mod_gofmt_invalid.txt
│   │   │   │   ├── mod_gonoproxy.txt
│   │   │   │   ├── mod_gopkg_unstable.txt
│   │   │   │   ├── mod_graph.txt
│   │   │   │   ├── mod_help.txt
│   │   │   │   ├── mod_import.txt
│   │   │   │   ├── mod_import_mod.txt
│   │   │   │   ├── mod_import_v1suffix.txt
│   │   │   │   ├── mod_indirect.txt
│   │   │   │   ├── mod_indirect_main.txt
│   │   │   │   ├── mod_indirect_tidy.txt
│   │   │   │   ├── mod_init_dep.txt
│   │   │   │   ├── mod_init_empty.txt
│   │   │   │   ├── mod_init_glide.txt
│   │   │   │   ├── mod_init_path.txt
│   │   │   │   ├── mod_install_versioned.txt
│   │   │   │   ├── mod_internal.txt
│   │   │   │   ├── mod_invalid_version.txt
│   │   │   │   ├── mod_list.txt
│   │   │   │   ├── mod_list_bad_import.txt
│   │   │   │   ├── mod_list_compiled_concurrent.txt
│   │   │   │   ├── mod_list_dir.txt
│   │   │   │   ├── mod_list_direct.txt
│   │   │   │   ├── mod_list_pseudo.txt
│   │   │   │   ├── mod_list_replace_dir.txt
│   │   │   │   ├── mod_list_std.txt
│   │   │   │   ├── mod_list_test.txt
│   │   │   │   ├── mod_list_upgrade.txt
│   │   │   │   ├── mod_list_upgrade_pseudo.txt
│   │   │   │   ├── mod_load_badchain.txt
│   │   │   │   ├── mod_load_badmod.txt
│   │   │   │   ├── mod_load_badzip.txt
│   │   │   │   ├── mod_local_replace.txt
│   │   │   │   ├── mod_missingpkg_prerelease.txt
│   │   │   │   ├── mod_modinfo.txt
│   │   │   │   ├── mod_multirepo.txt
│   │   │   │   ├── mod_nomod.txt
│   │   │   │   ├── mod_off.txt
│   │   │   │   ├── mod_off_init.txt
│   │   │   │   ├── mod_outside.txt
│   │   │   │   ├── mod_patterns.txt
│   │   │   │   ├── mod_patterns_vendor.txt
│   │   │   │   ├── mod_proxy_https.txt
│   │   │   │   ├── mod_proxy_list.txt
│   │   │   │   ├── mod_pseudo_cache.txt
│   │   │   │   ├── mod_query.txt
│   │   │   │   ├── mod_query_empty.txt
│   │   │   │   ├── mod_query_exclude.txt
│   │   │   │   ├── mod_readonly.txt
│   │   │   │   ├── mod_replace.txt
│   │   │   │   ├── mod_replace_import.txt
│   │   │   │   ├── mod_require_exclude.txt
│   │   │   │   ├── mod_run_internal.txt
│   │   │   │   ├── mod_run_path.txt
│   │   │   │   ├── mod_std_vendor.txt
│   │   │   │   ├── mod_string_alias.txt
│   │   │   │   ├── mod_sum_replaced.txt
│   │   │   │   ├── mod_sumdb.txt
│   │   │   │   ├── mod_sumdb_cache.txt
│   │   │   │   ├── mod_sumdb_file_path.txt
│   │   │   │   ├── mod_sumdb_golang.txt
│   │   │   │   ├── mod_sumdb_proxy.txt
│   │   │   │   ├── mod_symlink.txt
│   │   │   │   ├── mod_test.txt
│   │   │   │   ├── mod_test_cached.txt
│   │   │   │   ├── mod_test_files.txt
│   │   │   │   ├── mod_tidy.txt
│   │   │   │   ├── mod_tidy_error.txt
│   │   │   │   ├── mod_tidy_quote.txt
│   │   │   │   ├── mod_tidy_replace.txt
│   │   │   │   ├── mod_tidy_sum.txt
│   │   │   │   ├── mod_upgrade_patch.txt
│   │   │   │   ├── mod_vcs_missing.txt
│   │   │   │   ├── mod_vendor.txt
│   │   │   │   ├── mod_vendor_build.txt
│   │   │   │   ├── mod_vendor_nodeps.txt
│   │   │   │   ├── mod_vendor_replace.txt
│   │   │   │   ├── mod_verify.txt
│   │   │   │   ├── mod_versions.txt
│   │   │   │   ├── mod_why.txt
│   │   │   │   ├── pattern_syntax_error.txt
│   │   │   │   ├── prevent_sys_unix_import.txt
│   │   │   │   ├── run_hello.txt
│   │   │   │   ├── run_set_executable_name.txt
│   │   │   │   ├── run_wildcard.txt
│   │   │   │   ├── script_wait.txt
│   │   │   │   ├── std_vendor.txt
│   │   │   │   ├── sum_readonly.txt
│   │   │   │   ├── test_badtest.txt
│   │   │   │   ├── test_compile_binary.txt
│   │   │   │   ├── test_devnull.txt
│   │   │   │   ├── test_generated_main.txt
│   │   │   │   ├── test_go111module_cache.txt
│   │   │   │   ├── test_status.txt
│   │   │   │   ├── test_timeout.txt
│   │   │   │   ├── vendor_complex.txt
│   │   │   │   ├── version.txt
│   │   │   │   ├── vet_asm.txt
│   │   │   │   └── vet_deps.txt
│   │   │   ├── shadow
│   │   │   │   ├── root1
│   │   │   │   │   └── src
│   │   │   │   │       ├── foo
│   │   │   │   │       │   └── foo.go
│   │   │   │   │       └── math
│   │   │   │   │           └── math.go
│   │   │   │   └── root2
│   │   │   │       └── src
│   │   │   │           └── foo
│   │   │   │               └── foo.go
│   │   │   ├── src
│   │   │   │   ├── badc
│   │   │   │   │   ├── x.c
│   │   │   │   │   └── x.go
│   │   │   │   ├── badpkg
│   │   │   │   │   └── x.go
│   │   │   │   ├── bench
│   │   │   │   │   └── x_test.go
│   │   │   │   ├── benchfatal
│   │   │   │   │   └── x_test.go
│   │   │   │   ├── canonical
│   │   │   │   │   ├── a
│   │   │   │   │   │   ├── a.go
│   │   │   │   │   │   └── vendor
│   │   │   │   │   │       └── c
│   │   │   │   │   │           └── c.go
│   │   │   │   │   ├── b
│   │   │   │   │   │   └── b.go
│   │   │   │   │   └── d
│   │   │   │   │       └── d.go
│   │   │   │   ├── cgoasm
│   │   │   │   │   ├── p.go
│   │   │   │   │   └── p.s
│   │   │   │   ├── cgocover
│   │   │   │   │   ├── p.go
│   │   │   │   │   └── p_test.go
│   │   │   │   ├── cgocover2
│   │   │   │   │   ├── p.go
│   │   │   │   │   └── x_test.go
│   │   │   │   ├── cgocover3
│   │   │   │   │   ├── p.go
│   │   │   │   │   ├── p_test.go
│   │   │   │   │   └── x_test.go
│   │   │   │   ├── cgocover4
│   │   │   │   │   ├── notcgo.go
│   │   │   │   │   ├── p.go
│   │   │   │   │   └── x_test.go
│   │   │   │   ├── cgotest
│   │   │   │   │   └── m.go
│   │   │   │   ├── coverasm
│   │   │   │   │   ├── p.go
│   │   │   │   │   ├── p.s
│   │   │   │   │   └── p_test.go
│   │   │   │   ├── coverbad
│   │   │   │   │   ├── p.go
│   │   │   │   │   ├── p1.go
│   │   │   │   │   └── p_test.go
│   │   │   │   ├── coverdep
│   │   │   │   │   ├── p.go
│   │   │   │   │   ├── p1
│   │   │   │   │   │   └── p1.go
│   │   │   │   │   └── p_test.go
│   │   │   │   ├── coverdep2
│   │   │   │   │   ├── p1
│   │   │   │   │   │   ├── p.go
│   │   │   │   │   │   └── p_test.go
│   │   │   │   │   └── p2
│   │   │   │   │       └── p2.go
│   │   │   │   ├── coverdot1
│   │   │   │   │   └── p.go
│   │   │   │   ├── coverdot2
│   │   │   │   │   ├── p.go
│   │   │   │   │   └── p_test.go
│   │   │   │   ├── dupload
│   │   │   │   │   ├── dupload.go
│   │   │   │   │   ├── p
│   │   │   │   │   │   └── p.go
│   │   │   │   │   ├── p2
│   │   │   │   │   │   └── p2.go
│   │   │   │   │   └── vendor
│   │   │   │   │       └── p
│   │   │   │   │           └── p.go
│   │   │   │   ├── empty
│   │   │   │   │   ├── pkg
│   │   │   │   │   │   └── pkg.go
│   │   │   │   │   ├── pkgtest
│   │   │   │   │   │   ├── pkg.go
│   │   │   │   │   │   └── test_test.go
│   │   │   │   │   ├── pkgtestxtest
│   │   │   │   │   │   ├── pkg.go
│   │   │   │   │   │   ├── test_test.go
│   │   │   │   │   │   └── xtest_test.go
│   │   │   │   │   ├── pkgxtest
│   │   │   │   │   │   ├── pkg.go
│   │   │   │   │   │   └── xtest_test.go
│   │   │   │   │   ├── test
│   │   │   │   │   │   └── test_test.go
│   │   │   │   │   ├── testxtest
│   │   │   │   │   │   ├── test_test.go
│   │   │   │   │   │   └── xtest_test.go
│   │   │   │   │   └── xtest
│   │   │   │   │       └── xtest_test.go
│   │   │   │   ├── exclude
│   │   │   │   │   ├── empty
│   │   │   │   │   │   └── x.txt
│   │   │   │   │   ├── ignore
│   │   │   │   │   │   └── _x.go
│   │   │   │   │   ├── x.go
│   │   │   │   │   └── x_linux.go
│   │   │   │   ├── failfast_test.go
│   │   │   │   ├── gencycle
│   │   │   │   │   └── gencycle.go
│   │   │   │   ├── go-cmd-test
│   │   │   │   │   └── helloworld.go
│   │   │   │   ├── hello
│   │   │   │   │   └── hello.go
│   │   │   │   ├── importmain
│   │   │   │   │   ├── ismain
│   │   │   │   │   │   └── main.go
│   │   │   │   │   └── test
│   │   │   │   │       ├── test.go
│   │   │   │   │       └── test_test.go
│   │   │   │   ├── main_test
│   │   │   │   │   ├── m.go
│   │   │   │   │   └── m_test.go
│   │   │   │   ├── multimain
│   │   │   │   │   └── multimain_test.go
│   │   │   │   ├── my.pkg
│   │   │   │   │   ├── main
│   │   │   │   │   │   └── main.go
│   │   │   │   │   └── pkg.go
│   │   │   │   ├── not_main
│   │   │   │   │   └── not_main.go
│   │   │   │   ├── notest
│   │   │   │   │   └── hello.go
│   │   │   │   ├── run
│   │   │   │   │   ├── bad.go
│   │   │   │   │   ├── good.go
│   │   │   │   │   ├── internal
│   │   │   │   │   │   └── internal.go
│   │   │   │   │   └── subdir
│   │   │   │   │       └── internal
│   │   │   │   │           └── private
│   │   │   │   │               └── private.go
│   │   │   │   ├── skipper
│   │   │   │   │   └── skip_test.go
│   │   │   │   ├── sleepy1
│   │   │   │   │   └── p_test.go
│   │   │   │   ├── sleepy2
│   │   │   │   │   └── p_test.go
│   │   │   │   ├── sleepybad
│   │   │   │   │   └── p.go
│   │   │   │   ├── syntaxerror
│   │   │   │   │   ├── x.go
│   │   │   │   │   └── x_test.go
│   │   │   │   ├── testcache
│   │   │   │   │   └── testcache_test.go
│   │   │   │   ├── testcycle
│   │   │   │   │   ├── p1
│   │   │   │   │   │   ├── p1.go
│   │   │   │   │   │   └── p1_test.go
│   │   │   │   │   ├── p2
│   │   │   │   │   │   └── p2.go
│   │   │   │   │   ├── p3
│   │   │   │   │   │   ├── p3.go
│   │   │   │   │   │   └── p3_test.go
│   │   │   │   │   └── q1
│   │   │   │   │       ├── q1.go
│   │   │   │   │       └── q1_test.go
│   │   │   │   ├── testdep
│   │   │   │   │   ├── p1
│   │   │   │   │   │   ├── p1.go
│   │   │   │   │   │   └── p1_test.go
│   │   │   │   │   ├── p2
│   │   │   │   │   │   └── p2.go
│   │   │   │   │   └── p3
│   │   │   │   │       └── p3.go
│   │   │   │   ├── testlist
│   │   │   │   │   ├── bench_test.go
│   │   │   │   │   ├── example_test.go
│   │   │   │   │   └── test_test.go
│   │   │   │   ├── testnorun
│   │   │   │   │   └── p.go
│   │   │   │   ├── testrace
│   │   │   │   │   └── race_test.go
│   │   │   │   ├── testregexp
│   │   │   │   │   ├── x_test.go
│   │   │   │   │   └── z_test.go
│   │   │   │   ├── vend
│   │   │   │   │   ├── bad.go
│   │   │   │   │   ├── dir1
│   │   │   │   │   │   └── dir1.go
│   │   │   │   │   ├── good.go
│   │   │   │   │   ├── hello
│   │   │   │   │   │   ├── hello.go
│   │   │   │   │   │   ├── hello_test.go
│   │   │   │   │   │   └── hellox_test.go
│   │   │   │   │   ├── subdir
│   │   │   │   │   │   ├── bad.go
│   │   │   │   │   │   └── good.go
│   │   │   │   │   ├── vendor
│   │   │   │   │   │   ├── p
│   │   │   │   │   │   │   └── p.go
│   │   │   │   │   │   ├── q
│   │   │   │   │   │   │   └── q.go
│   │   │   │   │   │   ├── strings
│   │   │   │   │   │   │   └── msg.go
│   │   │   │   │   │   └── vend
│   │   │   │   │   │       └── dir1
│   │   │   │   │   │           └── dir2
│   │   │   │   │   │               └── dir2.go
│   │   │   │   │   └── x
│   │   │   │   │       ├── invalid
│   │   │   │   │       │   └── invalid.go
│   │   │   │   │       ├── vendor
│   │   │   │   │       │   ├── p
│   │   │   │   │       │   │   ├── p
│   │   │   │   │       │   │   │   └── p.go
│   │   │   │   │       │   │   └── p.go
│   │   │   │   │       │   └── r
│   │   │   │   │       │       └── r.go
│   │   │   │   │       └── x.go
│   │   │   │   ├── vetcycle
│   │   │   │   │   └── p.go
│   │   │   │   ├── vetfail
│   │   │   │   │   ├── p1
│   │   │   │   │   │   └── p1.go
│   │   │   │   │   └── p2
│   │   │   │   │       ├── p2.go
│   │   │   │   │       └── p2_test.go
│   │   │   │   ├── vetpkg
│   │   │   │   │   ├── a_test.go
│   │   │   │   │   ├── b.go
│   │   │   │   │   └── c.go
│   │   │   │   └── xtestonly
│   │   │   │       ├── f.go
│   │   │   │       └── f_test.go
│   │   │   ├── standalone_benchmark_test.go
│   │   │   ├── standalone_fail_sub_test.go
│   │   │   ├── standalone_main_normal_test.go
│   │   │   ├── standalone_main_wrong_test.go
│   │   │   ├── standalone_parallel_sub_test.go
│   │   │   ├── standalone_sub_test.go
│   │   │   ├── standalone_test.go
│   │   │   ├── standalone_testmain_flag_test.go
│   │   │   ├── testcover
│   │   │   │   ├── pkg1
│   │   │   │   │   └── a.go
│   │   │   │   ├── pkg2
│   │   │   │   │   ├── a.go
│   │   │   │   │   └── a_test.go
│   │   │   │   ├── pkg3
│   │   │   │   │   ├── a.go
│   │   │   │   │   └── a_test.go
│   │   │   │   └── pkg4
│   │   │   │       ├── a.go
│   │   │   │       └── a_test.go
│   │   │   ├── testimport
│   │   │   │   ├── p.go
│   │   │   │   ├── p1
│   │   │   │   │   └── p1.go
│   │   │   │   ├── p2
│   │   │   │   │   └── p2.go
│   │   │   │   ├── p_test.go
│   │   │   │   └── x_test.go
│   │   │   ├── testinternal
│   │   │   │   └── p.go
│   │   │   ├── testinternal2
│   │   │   │   ├── p.go
│   │   │   │   └── x
│   │   │   │       └── y
│   │   │   │           └── z
│   │   │   │               └── internal
│   │   │   │                   └── w
│   │   │   │                       └── w.go
│   │   │   ├── testinternal3
│   │   │   │   └── t.go
│   │   │   ├── testinternal4
│   │   │   │   └── src
│   │   │   │       ├── p
│   │   │   │       │   └── p.go
│   │   │   │       └── q
│   │   │   │           ├── internal
│   │   │   │           │   └── x
│   │   │   │           │       └── x.go
│   │   │   │           └── j
│   │   │   │               └── j.go
│   │   │   ├── testonly
│   │   │   │   └── p_test.go
│   │   │   ├── testonly2
│   │   │   │   └── t.go
│   │   │   ├── testterminal18153
│   │   │   │   └── terminal_test.go
│   │   │   ├── testvendor
│   │   │   │   └── src
│   │   │   │       ├── p
│   │   │   │       │   └── p.go
│   │   │   │       └── q
│   │   │   │           ├── vendor
│   │   │   │           │   └── x
│   │   │   │           │       └── x.go
│   │   │   │           ├── y
│   │   │   │           │   └── y.go
│   │   │   │           └── z
│   │   │   │               └── z.go
│   │   │   ├── testvendor2
│   │   │   │   ├── src
│   │   │   │   │   └── p
│   │   │   │   │       └── p.go
│   │   │   │   └── vendor
│   │   │   │       └── x
│   │   │   │           └── x.go
│   │   │   ├── timeoutbench_test.go
│   │   │   └── vendormod.txt
│   │   └── vendor_test.go
│   ├── go.mod
│   ├── go.sum
│   ├── gofmt
│   │   ├── doc.go
│   │   ├── gofmt.go
│   │   ├── gofmt_test.go
│   │   ├── internal.go
│   │   ├── long_test.go
│   │   ├── rewrite.go
│   │   ├── simplify.go
│   │   └── testdata
│   │       ├── comments.golden
│   │       ├── comments.input
│   │       ├── composites.golden
│   │       ├── composites.input
│   │       ├── crlf.golden
│   │       ├── crlf.input
│   │       ├── emptydecl.golden
│   │       ├── emptydecl.input
│   │       ├── go2numbers.golden
│   │       ├── go2numbers.input
│   │       ├── import.golden
│   │       ├── import.input
│   │       ├── ranges.golden
│   │       ├── ranges.input
│   │       ├── rewrite1.golden
│   │       ├── rewrite1.input
│   │       ├── rewrite2.golden
│   │       ├── rewrite2.input
│   │       ├── rewrite3.golden
│   │       ├── rewrite3.input
│   │       ├── rewrite4.golden
│   │       ├── rewrite4.input
│   │       ├── rewrite5.golden
│   │       ├── rewrite5.input
│   │       ├── rewrite6.golden
│   │       ├── rewrite6.input
│   │       ├── rewrite7.golden
│   │       ├── rewrite7.input
│   │       ├── rewrite8.golden
│   │       ├── rewrite8.input
│   │       ├── rewrite9.golden
│   │       ├── rewrite9.input
│   │       ├── slices1.golden
│   │       ├── slices1.input
│   │       ├── stdin1.golden
│   │       ├── stdin1.input
│   │       ├── stdin2.golden
│   │       ├── stdin2.input
│   │       ├── stdin3.golden
│   │       ├── stdin3.input
│   │       ├── stdin4.golden
│   │       ├── stdin4.input
│   │       ├── stdin5.golden
│   │       ├── stdin5.input
│   │       ├── stdin6.golden
│   │       ├── stdin6.input
│   │       ├── stdin7.golden
│   │       ├── stdin7.input
│   │       ├── typealias.golden
│   │       ├── typealias.input
│   │       ├── typeswitch.golden
│   │       └── typeswitch.input
│   ├── internal
│   │   ├── bio
│   │   │   ├── buf.go
│   │   │   ├── buf_mmap.go
│   │   │   ├── buf_nommap.go
│   │   │   └── must.go
│   │   ├── browser
│   │   │   └── browser.go
│   │   ├── buildid
│   │   │   ├── buildid.go
│   │   │   ├── buildid_test.go
│   │   │   ├── note.go
│   │   │   ├── rewrite.go
│   │   │   └── testdata
│   │   │       ├── a.elf
│   │   │       ├── a.macho
│   │   │       ├── a.pe
│   │   │       └── p.a
│   │   ├── dwarf
│   │   │   ├── dwarf.go
│   │   │   ├── dwarf_defs.go
│   │   │   └── dwarf_test.go
│   │   ├── edit
│   │   │   ├── edit.go
│   │   │   └── edit_test.go
│   │   ├── gcprog
│   │   │   └── gcprog.go
│   │   ├── goobj
│   │   │   ├── goobj_test.go
│   │   │   ├── read.go
│   │   │   └── testdata
│   │   │       ├── go1.go
│   │   │       ├── go2.go
│   │   │       └── mycgo
│   │   │           ├── c1.c
│   │   │           ├── c2.c
│   │   │           ├── go.go
│   │   │           ├── go1.go
│   │   │           └── go2.go
│   │   ├── obj
│   │   │   ├── abi_string.go
│   │   │   ├── addrtype_string.go
│   │   │   ├── arm
│   │   │   │   ├── a.out.go
│   │   │   │   ├── anames.go
│   │   │   │   ├── anames5.go
│   │   │   │   ├── asm5.go
│   │   │   │   ├── list5.go
│   │   │   │   └── obj5.go
│   │   │   ├── arm64
│   │   │   │   ├── a.out.go
│   │   │   │   ├── anames.go
│   │   │   │   ├── anames7.go
│   │   │   │   ├── asm7.go
│   │   │   │   ├── asm_test.go
│   │   │   │   ├── doc.go
│   │   │   │   ├── list7.go
│   │   │   │   └── obj7.go
│   │   │   ├── data.go
│   │   │   ├── go.go
│   │   │   ├── inl.go
│   │   │   ├── ld.go
│   │   │   ├── line.go
│   │   │   ├── line_test.go
│   │   │   ├── link.go
│   │   │   ├── mips
│   │   │   │   ├── a.out.go
│   │   │   │   ├── anames.go
│   │   │   │   ├── anames0.go
│   │   │   │   ├── asm0.go
│   │   │   │   ├── list0.go
│   │   │   │   └── obj0.go
│   │   │   ├── objfile.go
│   │   │   ├── pass.go
│   │   │   ├── pcln.go
│   │   │   ├── plist.go
│   │   │   ├── ppc64
│   │   │   │   ├── a.out.go
│   │   │   │   ├── anames.go
│   │   │   │   ├── anames9.go
│   │   │   │   ├── asm9.go
│   │   │   │   ├── doc.go
│   │   │   │   ├── list9.go
│   │   │   │   └── obj9.go
│   │   │   ├── s390x
│   │   │   │   ├── a.out.go
│   │   │   │   ├── anames.go
│   │   │   │   ├── anamesz.go
│   │   │   │   ├── asmz.go
│   │   │   │   ├── listz.go
│   │   │   │   ├── objz.go
│   │   │   │   └── vector.go
│   │   │   ├── sizeof_test.go
│   │   │   ├── stringer.go
│   │   │   ├── sym.go
│   │   │   ├── textflag.go
│   │   │   ├── util.go
│   │   │   ├── wasm
│   │   │   │   ├── a.out.go
│   │   │   │   ├── anames.go
│   │   │   │   └── wasmobj.go
│   │   │   └── x86
│   │   │       ├── a.out.go
│   │   │       ├── aenum.go
│   │   │       ├── anames.go
│   │   │       ├── asm6.go
│   │   │       ├── asm_test.go
│   │   │       ├── avx_optabs.go
│   │   │       ├── evex.go
│   │   │       ├── list6.go
│   │   │       ├── obj6.go
│   │   │       ├── obj6_test.go
│   │   │       ├── pcrelative_test.go
│   │   │       └── ytab.go
│   │   ├── objabi
│   │   │   ├── autotype.go
│   │   │   ├── doc.go
│   │   │   ├── flag.go
│   │   │   ├── funcdata.go
│   │   │   ├── funcid.go
│   │   │   ├── head.go
│   │   │   ├── line.go
│   │   │   ├── line_test.go
│   │   │   ├── path.go
│   │   │   ├── path_test.go
│   │   │   ├── reloctype.go
│   │   │   ├── reloctype_string.go
│   │   │   ├── stack.go
│   │   │   ├── symkind.go
│   │   │   ├── symkind_string.go
│   │   │   ├── typekind.go
│   │   │   ├── util.go
│   │   │   └── zbootstrap.go
│   │   ├── objfile
│   │   │   ├── disasm.go
│   │   │   ├── elf.go
│   │   │   ├── goobj.go
│   │   │   ├── macho.go
│   │   │   ├── objfile.go
│   │   │   ├── pe.go
│   │   │   ├── plan9obj.go
│   │   │   └── xcoff.go
│   │   ├── src
│   │   │   ├── pos.go
│   │   │   ├── pos_test.go
│   │   │   ├── xpos.go
│   │   │   └── xpos_test.go
│   │   ├── sys
│   │   │   ├── arch.go
│   │   │   └── supported.go
│   │   └── test2json
│   │       ├── test2json.go
│   │       ├── test2json_test.go
│   │       └── testdata
│   │           ├── ascii.json
│   │           ├── ascii.test
│   │           ├── bench.json
│   │           ├── bench.test
│   │           ├── benchfail.json
│   │           ├── benchfail.test
│   │           ├── benchshort.json
│   │           ├── benchshort.test
│   │           ├── issue23036.json
│   │           ├── issue23036.test
│   │           ├── issue23920.json
│   │           ├── issue23920.test
│   │           ├── smiley.json
│   │           ├── smiley.test
│   │           ├── unicode.json
│   │           ├── unicode.test
│   │           ├── vet.json
│   │           └── vet.test
│   ├── link
│   │   ├── doc.go
│   │   ├── dwarf_test.go
│   │   ├── elf_test.go
│   │   ├── internal
│   │   │   ├── amd64
│   │   │   │   ├── asm.go
│   │   │   │   ├── l.go
│   │   │   │   └── obj.go
│   │   │   ├── arm
│   │   │   │   ├── asm.go
│   │   │   │   ├── l.go
│   │   │   │   └── obj.go
│   │   │   ├── arm64
│   │   │   │   ├── asm.go
│   │   │   │   ├── l.go
│   │   │   │   └── obj.go
│   │   │   ├── ld
│   │   │   │   ├── ar.go
│   │   │   │   ├── config.go
│   │   │   │   ├── data.go
│   │   │   │   ├── deadcode.go
│   │   │   │   ├── decodesym.go
│   │   │   │   ├── dwarf.go
│   │   │   │   ├── dwarf_test.go
│   │   │   │   ├── elf.go
│   │   │   │   ├── go.go
│   │   │   │   ├── ld.go
│   │   │   │   ├── ld_test.go
│   │   │   │   ├── lib.go
│   │   │   │   ├── link.go
│   │   │   │   ├── macho.go
│   │   │   │   ├── macho_combine_dwarf.go
│   │   │   │   ├── main.go
│   │   │   │   ├── nooptcgolink_test.go
│   │   │   │   ├── outbuf.go
│   │   │   │   ├── outbuf_mmap.go
│   │   │   │   ├── outbuf_nommap.go
│   │   │   │   ├── pcln.go
│   │   │   │   ├── pe.go
│   │   │   │   ├── sym.go
│   │   │   │   ├── symtab.go
│   │   │   │   ├── testdata
│   │   │   │   │   ├── httptest
│   │   │   │   │   │   └── main
│   │   │   │   │   │       └── main.go
│   │   │   │   │   ├── issue10978
│   │   │   │   │   │   ├── main.go
│   │   │   │   │   │   └── main.s
│   │   │   │   │   ├── issue25459
│   │   │   │   │   │   ├── a
│   │   │   │   │   │   │   └── a.go
│   │   │   │   │   │   └── main
│   │   │   │   │   │       └── main.go
│   │   │   │   │   ├── issue26237
│   │   │   │   │   │   ├── b.dir
│   │   │   │   │   │   │   └── b.go
│   │   │   │   │   │   └── main
│   │   │   │   │   │       └── main.go
│   │   │   │   │   └── issue32233
│   │   │   │   │       ├── lib
│   │   │   │   │       │   ├── ObjC.m
│   │   │   │   │       │   └── lib.go
│   │   │   │   │       └── main
│   │   │   │   │           └── main.go
│   │   │   │   ├── typelink.go
│   │   │   │   ├── util.go
│   │   │   │   └── xcoff.go
│   │   │   ├── loadelf
│   │   │   │   └── ldelf.go
│   │   │   ├── loadmacho
│   │   │   │   └── ldmacho.go
│   │   │   ├── loadpe
│   │   │   │   └── ldpe.go
│   │   │   ├── loadxcoff
│   │   │   │   └── ldxcoff.go
│   │   │   ├── mips
│   │   │   │   ├── asm.go
│   │   │   │   ├── l.go
│   │   │   │   └── obj.go
│   │   │   ├── mips64
│   │   │   │   ├── asm.go
│   │   │   │   ├── l.go
│   │   │   │   └── obj.go
│   │   │   ├── objfile
│   │   │   │   └── objfile.go
│   │   │   ├── ppc64
│   │   │   │   ├── asm.go
│   │   │   │   ├── l.go
│   │   │   │   └── obj.go
│   │   │   ├── s390x
│   │   │   │   ├── asm.go
│   │   │   │   ├── l.go
│   │   │   │   └── obj.go
│   │   │   ├── sym
│   │   │   │   ├── attribute.go
│   │   │   │   ├── library.go
│   │   │   │   ├── reloc.go
│   │   │   │   ├── segment.go
│   │   │   │   ├── sizeof_test.go
│   │   │   │   ├── symbol.go
│   │   │   │   ├── symbols.go
│   │   │   │   ├── symkind.go
│   │   │   │   └── symkind_string.go
│   │   │   ├── wasm
│   │   │   │   ├── asm.go
│   │   │   │   └── obj.go
│   │   │   └── x86
│   │   │       ├── asm.go
│   │   │       ├── l.go
│   │   │       └── obj.go
│   │   ├── link_test.go
│   │   ├── linkbig_test.go
│   │   ├── main.go
│   │   └── testdata
│   │       ├── lib.go
│   │       └── main.m
│   ├── nm
│   │   ├── doc.go
│   │   ├── nm.go
│   │   ├── nm_cgo_test.go
│   │   └── nm_test.go
│   ├── objdump
│   │   ├── main.go
│   │   ├── objdump_test.go
│   │   └── testdata
│   │       └── fmthello.go
│   ├── pack
│   │   ├── doc.go
│   │   ├── pack.go
│   │   └── pack_test.go
│   ├── pprof
│   │   ├── README
│   │   ├── doc.go
│   │   ├── pprof.go
│   │   └── readlineui.go
│   ├── test2json
│   │   └── main.go
│   ├── trace
│   │   ├── annotations.go
│   │   ├── annotations_test.go
│   │   ├── doc.go
│   │   ├── goroutines.go
│   │   ├── main.go
│   │   ├── mmu.go
│   │   ├── pprof.go
│   │   ├── trace.go
│   │   ├── trace_test.go
│   │   └── trace_unix_test.go
│   ├── vendor
│   │   ├── github.com
│   │   │   ├── google
│   │   │   │   └── pprof
│   │   │   │       ├── AUTHORS
│   │   │   │       ├── CONTRIBUTORS
│   │   │   │       ├── LICENSE
│   │   │   │       ├── driver
│   │   │   │       │   └── driver.go
│   │   │   │       ├── internal
│   │   │   │       │   ├── binutils
│   │   │   │       │   │   ├── addr2liner.go
│   │   │   │       │   │   ├── addr2liner_llvm.go
│   │   │   │       │   │   ├── addr2liner_nm.go
│   │   │   │       │   │   ├── binutils.go
│   │   │   │       │   │   └── disasm.go
│   │   │   │       │   ├── driver
│   │   │   │       │   │   ├── cli.go
│   │   │   │       │   │   ├── commands.go
│   │   │   │       │   │   ├── driver.go
│   │   │   │       │   │   ├── driver_focus.go
│   │   │   │       │   │   ├── fetch.go
│   │   │   │       │   │   ├── flags.go
│   │   │   │       │   │   ├── flamegraph.go
│   │   │   │       │   │   ├── interactive.go
│   │   │   │       │   │   ├── options.go
│   │   │   │       │   │   ├── svg.go
│   │   │   │       │   │   ├── tempfile.go
│   │   │   │       │   │   ├── webhtml.go
│   │   │   │       │   │   └── webui.go
│   │   │   │       │   ├── elfexec
│   │   │   │       │   │   └── elfexec.go
│   │   │   │       │   ├── graph
│   │   │   │       │   │   ├── dotgraph.go
│   │   │   │       │   │   └── graph.go
│   │   │   │       │   ├── measurement
│   │   │   │       │   │   └── measurement.go
│   │   │   │       │   ├── plugin
│   │   │   │       │   │   └── plugin.go
│   │   │   │       │   ├── report
│   │   │   │       │   │   ├── report.go
│   │   │   │       │   │   ├── source.go
│   │   │   │       │   │   └── source_html.go
│   │   │   │       │   ├── symbolizer
│   │   │   │       │   │   └── symbolizer.go
│   │   │   │       │   ├── symbolz
│   │   │   │       │   │   └── symbolz.go
│   │   │   │       │   └── transport
│   │   │   │       │       └── transport.go
│   │   │   │       ├── profile
│   │   │   │       │   ├── encode.go
│   │   │   │       │   ├── filter.go
│   │   │   │       │   ├── index.go
│   │   │   │       │   ├── legacy_java_profile.go
│   │   │   │       │   ├── legacy_profile.go
│   │   │   │       │   ├── merge.go
│   │   │   │       │   ├── profile.go
│   │   │   │       │   ├── proto.go
│   │   │   │       │   └── prune.go
│   │   │   │       └── third_party
│   │   │   │           ├── d3
│   │   │   │           │   ├── LICENSE
│   │   │   │           │   ├── README.md
│   │   │   │           │   └── d3.go
│   │   │   │           ├── d3flamegraph
│   │   │   │           │   ├── LICENSE
│   │   │   │           │   └── d3_flame_graph.go
│   │   │   │           └── svgpan
│   │   │   │               ├── LICENSE
│   │   │   │               └── svgpan.go
│   │   │   └── ianlancetaylor
│   │   │       └── demangle
│   │   │           ├── LICENSE
│   │   │           ├── README.md
│   │   │           ├── ast.go
│   │   │           └── demangle.go
│   │   ├── golang.org
│   │   │   └── x
│   │   │       ├── arch
│   │   │       │   ├── AUTHORS
│   │   │       │   ├── CONTRIBUTORS
│   │   │       │   ├── LICENSE
│   │   │       │   ├── PATENTS
│   │   │       │   ├── arm
│   │   │       │   │   └── armasm
│   │   │       │   │       ├── Makefile
│   │   │       │   │       ├── decode.go
│   │   │       │   │       ├── gnu.go
│   │   │       │   │       ├── inst.go
│   │   │       │   │       ├── plan9x.go
│   │   │       │   │       └── tables.go
│   │   │       │   ├── arm64
│   │   │       │   │   └── arm64asm
│   │   │       │   │       ├── arg.go
│   │   │       │   │       ├── condition.go
│   │   │       │   │       ├── condition_util.go
│   │   │       │   │       ├── decode.go
│   │   │       │   │       ├── gnu.go
│   │   │       │   │       ├── inst.go
│   │   │       │   │       ├── inst.json
│   │   │       │   │       ├── plan9x.go
│   │   │       │   │       └── tables.go
│   │   │       │   ├── ppc64
│   │   │       │   │   └── ppc64asm
│   │   │       │   │       ├── decode.go
│   │   │       │   │       ├── doc.go
│   │   │       │   │       ├── field.go
│   │   │       │   │       ├── gnu.go
│   │   │       │   │       ├── inst.go
│   │   │       │   │       ├── plan9.go
│   │   │       │   │       └── tables.go
│   │   │       │   └── x86
│   │   │       │       └── x86asm
│   │   │       │           ├── Makefile
│   │   │       │           ├── decode.go
│   │   │       │           ├── gnu.go
│   │   │       │           ├── inst.go
│   │   │       │           ├── intel.go
│   │   │       │           ├── plan9x.go
│   │   │       │           └── tables.go
│   │   │       ├── crypto
│   │   │       │   ├── AUTHORS
│   │   │       │   ├── CONTRIBUTORS
│   │   │       │   ├── LICENSE
│   │   │       │   ├── PATENTS
│   │   │       │   └── ssh
│   │   │       │       └── terminal
│   │   │       │           ├── terminal.go
│   │   │       │           ├── util.go
│   │   │       │           ├── util_aix.go
│   │   │       │           ├── util_bsd.go
│   │   │       │           ├── util_linux.go
│   │   │       │           ├── util_plan9.go
│   │   │       │           ├── util_solaris.go
│   │   │       │           └── util_windows.go
│   │   │       ├── sys
│   │   │       │   ├── AUTHORS
│   │   │       │   ├── CONTRIBUTORS
│   │   │       │   ├── LICENSE
│   │   │       │   ├── PATENTS
│   │   │       │   ├── unix
│   │   │       │   │   ├── README.md
│   │   │       │   │   ├── affinity_linux.go
│   │   │       │   │   ├── aliases.go
│   │   │       │   │   ├── asm_aix_ppc64.s
│   │   │       │   │   ├── asm_darwin_386.s
│   │   │       │   │   ├── asm_darwin_amd64.s
│   │   │       │   │   ├── asm_darwin_arm.s
│   │   │       │   │   ├── asm_darwin_arm64.s
│   │   │       │   │   ├── asm_dragonfly_amd64.s
│   │   │       │   │   ├── asm_freebsd_386.s
│   │   │       │   │   ├── asm_freebsd_amd64.s
│   │   │       │   │   ├── asm_freebsd_arm.s
│   │   │       │   │   ├── asm_freebsd_arm64.s
│   │   │       │   │   ├── asm_linux_386.s
│   │   │       │   │   ├── asm_linux_amd64.s
│   │   │       │   │   ├── asm_linux_arm.s
│   │   │       │   │   ├── asm_linux_arm64.s
│   │   │       │   │   ├── asm_linux_mips64x.s
│   │   │       │   │   ├── asm_linux_mipsx.s
│   │   │       │   │   ├── asm_linux_ppc64x.s
│   │   │       │   │   ├── asm_linux_s390x.s
│   │   │       │   │   ├── asm_netbsd_386.s
│   │   │       │   │   ├── asm_netbsd_amd64.s
│   │   │       │   │   ├── asm_netbsd_arm.s
│   │   │       │   │   ├── asm_netbsd_arm64.s
│   │   │       │   │   ├── asm_openbsd_386.s
│   │   │       │   │   ├── asm_openbsd_amd64.s
│   │   │       │   │   ├── asm_openbsd_arm.s
│   │   │       │   │   ├── asm_openbsd_arm64.s
│   │   │       │   │   ├── asm_solaris_amd64.s
│   │   │       │   │   ├── bluetooth_linux.go
│   │   │       │   │   ├── cap_freebsd.go
│   │   │       │   │   ├── constants.go
│   │   │       │   │   ├── dev_aix_ppc.go
│   │   │       │   │   ├── dev_aix_ppc64.go
│   │   │       │   │   ├── dev_darwin.go
│   │   │       │   │   ├── dev_dragonfly.go
│   │   │       │   │   ├── dev_freebsd.go
│   │   │       │   │   ├── dev_linux.go
│   │   │       │   │   ├── dev_netbsd.go
│   │   │       │   │   ├── dev_openbsd.go
│   │   │       │   │   ├── dirent.go
│   │   │       │   │   ├── endian_big.go
│   │   │       │   │   ├── endian_little.go
│   │   │       │   │   ├── env_unix.go
│   │   │       │   │   ├── errors_freebsd_386.go
│   │   │       │   │   ├── errors_freebsd_amd64.go
│   │   │       │   │   ├── errors_freebsd_arm.go
│   │   │       │   │   ├── fcntl.go
│   │   │       │   │   ├── fcntl_darwin.go
│   │   │       │   │   ├── fcntl_linux_32bit.go
│   │   │       │   │   ├── gccgo.go
│   │   │       │   │   ├── gccgo_c.c
│   │   │       │   │   ├── gccgo_linux_amd64.go
│   │   │       │   │   ├── ioctl.go
│   │   │       │   │   ├── mkall.sh
│   │   │       │   │   ├── mkerrors.sh
│   │   │       │   │   ├── pagesize_unix.go
│   │   │       │   │   ├── pledge_openbsd.go
│   │   │       │   │   ├── race.go
│   │   │       │   │   ├── race0.go
│   │   │       │   │   ├── sockcmsg_linux.go
│   │   │       │   │   ├── sockcmsg_unix.go
│   │   │       │   │   ├── str.go
│   │   │       │   │   ├── syscall.go
│   │   │       │   │   ├── syscall_aix.go
│   │   │       │   │   ├── syscall_aix_ppc.go
│   │   │       │   │   ├── syscall_aix_ppc64.go
│   │   │       │   │   ├── syscall_bsd.go
│   │   │       │   │   ├── syscall_darwin.go
│   │   │       │   │   ├── syscall_darwin_386.go
│   │   │       │   │   ├── syscall_darwin_amd64.go
│   │   │       │   │   ├── syscall_darwin_arm.go
│   │   │       │   │   ├── syscall_darwin_arm64.go
│   │   │       │   │   ├── syscall_darwin_libSystem.go
│   │   │       │   │   ├── syscall_dragonfly.go
│   │   │       │   │   ├── syscall_dragonfly_amd64.go
│   │   │       │   │   ├── syscall_freebsd.go
│   │   │       │   │   ├── syscall_freebsd_386.go
│   │   │       │   │   ├── syscall_freebsd_amd64.go
│   │   │       │   │   ├── syscall_freebsd_arm.go
│   │   │       │   │   ├── syscall_freebsd_arm64.go
│   │   │       │   │   ├── syscall_linux.go
│   │   │       │   │   ├── syscall_linux_386.go
│   │   │       │   │   ├── syscall_linux_amd64.go
│   │   │       │   │   ├── syscall_linux_amd64_gc.go
│   │   │       │   │   ├── syscall_linux_arm.go
│   │   │       │   │   ├── syscall_linux_arm64.go
│   │   │       │   │   ├── syscall_linux_gc.go
│   │   │       │   │   ├── syscall_linux_gc_386.go
│   │   │       │   │   ├── syscall_linux_gccgo_386.go
│   │   │       │   │   ├── syscall_linux_gccgo_arm.go
│   │   │       │   │   ├── syscall_linux_mips64x.go
│   │   │       │   │   ├── syscall_linux_mipsx.go
│   │   │       │   │   ├── syscall_linux_ppc64x.go
│   │   │       │   │   ├── syscall_linux_riscv64.go
│   │   │       │   │   ├── syscall_linux_s390x.go
│   │   │       │   │   ├── syscall_linux_sparc64.go
│   │   │       │   │   ├── syscall_netbsd.go
│   │   │       │   │   ├── syscall_netbsd_386.go
│   │   │       │   │   ├── syscall_netbsd_amd64.go
│   │   │       │   │   ├── syscall_netbsd_arm.go
│   │   │       │   │   ├── syscall_netbsd_arm64.go
│   │   │       │   │   ├── syscall_openbsd.go
│   │   │       │   │   ├── syscall_openbsd_386.go
│   │   │       │   │   ├── syscall_openbsd_amd64.go
│   │   │       │   │   ├── syscall_openbsd_arm.go
│   │   │       │   │   ├── syscall_openbsd_arm64.go
│   │   │       │   │   ├── syscall_solaris.go
│   │   │       │   │   ├── syscall_solaris_amd64.go
│   │   │       │   │   ├── syscall_unix.go
│   │   │       │   │   ├── syscall_unix_gc.go
│   │   │       │   │   ├── syscall_unix_gc_ppc64x.go
│   │   │       │   │   ├── timestruct.go
│   │   │       │   │   ├── unveil_openbsd.go
│   │   │       │   │   ├── xattr_bsd.go
│   │   │       │   │   ├── zerrors_aix_ppc.go
│   │   │       │   │   ├── zerrors_aix_ppc64.go
│   │   │       │   │   ├── zerrors_darwin_386.go
│   │   │       │   │   ├── zerrors_darwin_amd64.go
│   │   │       │   │   ├── zerrors_darwin_arm.go
│   │   │       │   │   ├── zerrors_darwin_arm64.go
│   │   │       │   │   ├── zerrors_dragonfly_amd64.go
│   │   │       │   │   ├── zerrors_freebsd_386.go
│   │   │       │   │   ├── zerrors_freebsd_amd64.go
│   │   │       │   │   ├── zerrors_freebsd_arm.go
│   │   │       │   │   ├── zerrors_freebsd_arm64.go
│   │   │       │   │   ├── zerrors_linux_386.go
│   │   │       │   │   ├── zerrors_linux_amd64.go
│   │   │       │   │   ├── zerrors_linux_arm.go
│   │   │       │   │   ├── zerrors_linux_arm64.go
│   │   │       │   │   ├── zerrors_linux_mips.go
│   │   │       │   │   ├── zerrors_linux_mips64.go
│   │   │       │   │   ├── zerrors_linux_mips64le.go
│   │   │       │   │   ├── zerrors_linux_mipsle.go
│   │   │       │   │   ├── zerrors_linux_ppc64.go
│   │   │       │   │   ├── zerrors_linux_ppc64le.go
│   │   │       │   │   ├── zerrors_linux_riscv64.go
│   │   │       │   │   ├── zerrors_linux_s390x.go
│   │   │       │   │   ├── zerrors_linux_sparc64.go
│   │   │       │   │   ├── zerrors_netbsd_386.go
│   │   │       │   │   ├── zerrors_netbsd_amd64.go
│   │   │       │   │   ├── zerrors_netbsd_arm.go
│   │   │       │   │   ├── zerrors_netbsd_arm64.go
│   │   │       │   │   ├── zerrors_openbsd_386.go
│   │   │       │   │   ├── zerrors_openbsd_amd64.go
│   │   │       │   │   ├── zerrors_openbsd_arm.go
│   │   │       │   │   ├── zerrors_openbsd_arm64.go
│   │   │       │   │   ├── zerrors_solaris_amd64.go
│   │   │       │   │   ├── zptrace386_linux.go
│   │   │       │   │   ├── zptracearm_linux.go
│   │   │       │   │   ├── zptracemips_linux.go
│   │   │       │   │   ├── zptracemipsle_linux.go
│   │   │       │   │   ├── zsyscall_aix_ppc.go
│   │   │       │   │   ├── zsyscall_aix_ppc64.go
│   │   │       │   │   ├── zsyscall_aix_ppc64_gc.go
│   │   │       │   │   ├── zsyscall_aix_ppc64_gccgo.go
│   │   │       │   │   ├── zsyscall_darwin_386.1_11.go
│   │   │       │   │   ├── zsyscall_darwin_386.go
│   │   │       │   │   ├── zsyscall_darwin_386.s
│   │   │       │   │   ├── zsyscall_darwin_amd64.1_11.go
│   │   │       │   │   ├── zsyscall_darwin_amd64.go
│   │   │       │   │   ├── zsyscall_darwin_amd64.s
│   │   │       │   │   ├── zsyscall_darwin_arm.1_11.go
│   │   │       │   │   ├── zsyscall_darwin_arm.go
│   │   │       │   │   ├── zsyscall_darwin_arm.s
│   │   │       │   │   ├── zsyscall_darwin_arm64.1_11.go
│   │   │       │   │   ├── zsyscall_darwin_arm64.go
│   │   │       │   │   ├── zsyscall_darwin_arm64.s
│   │   │       │   │   ├── zsyscall_dragonfly_amd64.go
│   │   │       │   │   ├── zsyscall_freebsd_386.go
│   │   │       │   │   ├── zsyscall_freebsd_amd64.go
│   │   │       │   │   ├── zsyscall_freebsd_arm.go
│   │   │       │   │   ├── zsyscall_freebsd_arm64.go
│   │   │       │   │   ├── zsyscall_linux_386.go
│   │   │       │   │   ├── zsyscall_linux_amd64.go
│   │   │       │   │   ├── zsyscall_linux_arm.go
│   │   │       │   │   ├── zsyscall_linux_arm64.go
│   │   │       │   │   ├── zsyscall_linux_mips.go
│   │   │       │   │   ├── zsyscall_linux_mips64.go
│   │   │       │   │   ├── zsyscall_linux_mips64le.go
│   │   │       │   │   ├── zsyscall_linux_mipsle.go
│   │   │       │   │   ├── zsyscall_linux_ppc64.go
│   │   │       │   │   ├── zsyscall_linux_ppc64le.go
│   │   │       │   │   ├── zsyscall_linux_riscv64.go
│   │   │       │   │   ├── zsyscall_linux_s390x.go
│   │   │       │   │   ├── zsyscall_linux_sparc64.go
│   │   │       │   │   ├── zsyscall_netbsd_386.go
│   │   │       │   │   ├── zsyscall_netbsd_amd64.go
│   │   │       │   │   ├── zsyscall_netbsd_arm.go
│   │   │       │   │   ├── zsyscall_netbsd_arm64.go
│   │   │       │   │   ├── zsyscall_openbsd_386.go
│   │   │       │   │   ├── zsyscall_openbsd_amd64.go
│   │   │       │   │   ├── zsyscall_openbsd_arm.go
│   │   │       │   │   ├── zsyscall_openbsd_arm64.go
│   │   │       │   │   ├── zsyscall_solaris_amd64.go
│   │   │       │   │   ├── zsysctl_openbsd_386.go
│   │   │       │   │   ├── zsysctl_openbsd_amd64.go
│   │   │       │   │   ├── zsysctl_openbsd_arm.go
│   │   │       │   │   ├── zsysctl_openbsd_arm64.go
│   │   │       │   │   ├── zsysnum_darwin_386.go
│   │   │       │   │   ├── zsysnum_darwin_amd64.go
│   │   │       │   │   ├── zsysnum_darwin_arm.go
│   │   │       │   │   ├── zsysnum_darwin_arm64.go
│   │   │       │   │   ├── zsysnum_dragonfly_amd64.go
│   │   │       │   │   ├── zsysnum_freebsd_386.go
│   │   │       │   │   ├── zsysnum_freebsd_amd64.go
│   │   │       │   │   ├── zsysnum_freebsd_arm.go
│   │   │       │   │   ├── zsysnum_freebsd_arm64.go
│   │   │       │   │   ├── zsysnum_linux_386.go
│   │   │       │   │   ├── zsysnum_linux_amd64.go
│   │   │       │   │   ├── zsysnum_linux_arm.go
│   │   │       │   │   ├── zsysnum_linux_arm64.go
│   │   │       │   │   ├── zsysnum_linux_mips.go
│   │   │       │   │   ├── zsysnum_linux_mips64.go
│   │   │       │   │   ├── zsysnum_linux_mips64le.go
│   │   │       │   │   ├── zsysnum_linux_mipsle.go
│   │   │       │   │   ├── zsysnum_linux_ppc64.go
│   │   │       │   │   ├── zsysnum_linux_ppc64le.go
│   │   │       │   │   ├── zsysnum_linux_riscv64.go
│   │   │       │   │   ├── zsysnum_linux_s390x.go
│   │   │       │   │   ├── zsysnum_linux_sparc64.go
│   │   │       │   │   ├── zsysnum_netbsd_386.go
│   │   │       │   │   ├── zsysnum_netbsd_amd64.go
│   │   │       │   │   ├── zsysnum_netbsd_arm.go
│   │   │       │   │   ├── zsysnum_netbsd_arm64.go
│   │   │       │   │   ├── zsysnum_openbsd_386.go
│   │   │       │   │   ├── zsysnum_openbsd_amd64.go
│   │   │       │   │   ├── zsysnum_openbsd_arm.go
│   │   │       │   │   ├── zsysnum_openbsd_arm64.go
│   │   │       │   │   ├── ztypes_aix_ppc.go
│   │   │       │   │   ├── ztypes_aix_ppc64.go
│   │   │       │   │   ├── ztypes_darwin_386.go
│   │   │       │   │   ├── ztypes_darwin_amd64.go
│   │   │       │   │   ├── ztypes_darwin_arm.go
│   │   │       │   │   ├── ztypes_darwin_arm64.go
│   │   │       │   │   ├── ztypes_dragonfly_amd64.go
│   │   │       │   │   ├── ztypes_freebsd_386.go
│   │   │       │   │   ├── ztypes_freebsd_amd64.go
│   │   │       │   │   ├── ztypes_freebsd_arm.go
│   │   │       │   │   ├── ztypes_freebsd_arm64.go
│   │   │       │   │   ├── ztypes_linux_386.go
│   │   │       │   │   ├── ztypes_linux_amd64.go
│   │   │       │   │   ├── ztypes_linux_arm.go
│   │   │       │   │   ├── ztypes_linux_arm64.go
│   │   │       │   │   ├── ztypes_linux_mips.go
│   │   │       │   │   ├── ztypes_linux_mips64.go
│   │   │       │   │   ├── ztypes_linux_mips64le.go
│   │   │       │   │   ├── ztypes_linux_mipsle.go
│   │   │       │   │   ├── ztypes_linux_ppc64.go
│   │   │       │   │   ├── ztypes_linux_ppc64le.go
│   │   │       │   │   ├── ztypes_linux_riscv64.go
│   │   │       │   │   ├── ztypes_linux_s390x.go
│   │   │       │   │   ├── ztypes_linux_sparc64.go
│   │   │       │   │   ├── ztypes_netbsd_386.go
│   │   │       │   │   ├── ztypes_netbsd_amd64.go
│   │   │       │   │   ├── ztypes_netbsd_arm.go
│   │   │       │   │   ├── ztypes_netbsd_arm64.go
│   │   │       │   │   ├── ztypes_openbsd_386.go
│   │   │       │   │   ├── ztypes_openbsd_amd64.go
│   │   │       │   │   ├── ztypes_openbsd_arm.go
│   │   │       │   │   ├── ztypes_openbsd_arm64.go
│   │   │       │   │   └── ztypes_solaris_amd64.go
│   │   │       │   └── windows
│   │   │       │       ├── aliases.go
│   │   │       │       ├── asm_windows_386.s
│   │   │       │       ├── asm_windows_amd64.s
│   │   │       │       ├── asm_windows_arm.s
│   │   │       │       ├── dll_windows.go
│   │   │       │       ├── env_windows.go
│   │   │       │       ├── eventlog.go
│   │   │       │       ├── exec_windows.go
│   │   │       │       ├── memory_windows.go
│   │   │       │       ├── mkerrors.bash
│   │   │       │       ├── mkerrors.go
│   │   │       │       ├── mksyscall.go
│   │   │       │       ├── race.go
│   │   │       │       ├── race0.go
│   │   │       │       ├── security_windows.go
│   │   │       │       ├── service.go
│   │   │       │       ├── str.go
│   │   │       │       ├── syscall.go
│   │   │       │       ├── syscall_windows.go
│   │   │       │       ├── types_windows.go
│   │   │       │       ├── types_windows_386.go
│   │   │       │       ├── types_windows_amd64.go
│   │   │       │       ├── types_windows_arm.go
│   │   │       │       ├── zerrors_windows.go
│   │   │       │       └── zsyscall_windows.go
│   │   │       └── tools
│   │   │           ├── AUTHORS
│   │   │           ├── CONTRIBUTORS
│   │   │           ├── LICENSE
│   │   │           ├── PATENTS
│   │   │           └── go
│   │   │               ├── analysis
│   │   │               │   ├── analysis.go
│   │   │               │   ├── doc.go
│   │   │               │   ├── internal
│   │   │               │   │   ├── analysisflags
│   │   │               │   │   │   ├── flags.go
│   │   │               │   │   │   └── help.go
│   │   │               │   │   └── facts
│   │   │               │   │       ├── facts.go
│   │   │               │   │       └── imports.go
│   │   │               │   ├── passes
│   │   │               │   │   ├── asmdecl
│   │   │               │   │   │   └── asmdecl.go
│   │   │               │   │   ├── assign
│   │   │               │   │   │   └── assign.go
│   │   │               │   │   ├── atomic
│   │   │               │   │   │   └── atomic.go
│   │   │               │   │   ├── bools
│   │   │               │   │   │   └── bools.go
│   │   │               │   │   ├── buildtag
│   │   │               │   │   │   └── buildtag.go
│   │   │               │   │   ├── cgocall
│   │   │               │   │   │   └── cgocall.go
│   │   │               │   │   ├── composite
│   │   │               │   │   │   ├── composite.go
│   │   │               │   │   │   └── whitelist.go
│   │   │               │   │   ├── copylock
│   │   │               │   │   │   └── copylock.go
│   │   │               │   │   ├── ctrlflow
│   │   │               │   │   │   └── ctrlflow.go
│   │   │               │   │   ├── errorsas
│   │   │               │   │   │   └── errorsas.go
│   │   │               │   │   ├── httpresponse
│   │   │               │   │   │   └── httpresponse.go
│   │   │               │   │   ├── inspect
│   │   │               │   │   │   └── inspect.go
│   │   │               │   │   ├── internal
│   │   │               │   │   │   └── analysisutil
│   │   │               │   │   │       └── util.go
│   │   │               │   │   ├── loopclosure
│   │   │               │   │   │   └── loopclosure.go
│   │   │               │   │   ├── lostcancel
│   │   │               │   │   │   └── lostcancel.go
│   │   │               │   │   ├── nilfunc
│   │   │               │   │   │   └── nilfunc.go
│   │   │               │   │   ├── printf
│   │   │               │   │   │   ├── printf.go
│   │   │               │   │   │   └── types.go
│   │   │               │   │   ├── shift
│   │   │               │   │   │   ├── dead.go
│   │   │               │   │   │   └── shift.go
│   │   │               │   │   ├── stdmethods
│   │   │               │   │   │   └── stdmethods.go
│   │   │               │   │   ├── structtag
│   │   │               │   │   │   └── structtag.go
│   │   │               │   │   ├── tests
│   │   │               │   │   │   └── tests.go
│   │   │               │   │   ├── unmarshal
│   │   │               │   │   │   └── unmarshal.go
│   │   │               │   │   ├── unreachable
│   │   │               │   │   │   └── unreachable.go
│   │   │               │   │   ├── unsafeptr
│   │   │               │   │   │   └── unsafeptr.go
│   │   │               │   │   └── unusedresult
│   │   │               │   │       └── unusedresult.go
│   │   │               │   ├── unitchecker
│   │   │               │   │   ├── unitchecker.go
│   │   │               │   │   └── unitchecker112.go
│   │   │               │   └── validate.go
│   │   │               ├── ast
│   │   │               │   ├── astutil
│   │   │               │   │   ├── enclosing.go
│   │   │               │   │   ├── imports.go
│   │   │               │   │   ├── rewrite.go
│   │   │               │   │   └── util.go
│   │   │               │   └── inspector
│   │   │               │       ├── inspector.go
│   │   │               │       └── typeof.go
│   │   │               ├── cfg
│   │   │               │   ├── builder.go
│   │   │               │   └── cfg.go
│   │   │               └── types
│   │   │                   ├── objectpath
│   │   │                   │   └── objectpath.go
│   │   │                   └── typeutil
│   │   │                       ├── callee.go
│   │   │                       ├── imports.go
│   │   │                       ├── map.go
│   │   │                       ├── methodsetcache.go
│   │   │                       └── ui.go
│   │   └── modules.txt
│   └── vet
│       ├── README
│       ├── doc.go
│       ├── main.go
│       ├── testdata
│       │   ├── asm
│       │   │   ├── asm.go
│       │   │   └── asm1.s
│       │   ├── assign
│       │   │   └── assign.go
│       │   ├── atomic
│       │   │   └── atomic.go
│       │   ├── bool
│       │   │   └── bool.go
│       │   ├── buildtag
│       │   │   └── buildtag.go
│       │   ├── cgo
│       │   │   └── cgo.go
│       │   ├── composite
│       │   │   └── composite.go
│       │   ├── copylock
│       │   │   └── copylock.go
│       │   ├── deadcode
│       │   │   └── deadcode.go
│       │   ├── httpresponse
│       │   │   └── httpresponse.go
│       │   ├── lostcancel
│       │   │   └── lostcancel.go
│       │   ├── method
│       │   │   └── method.go
│       │   ├── nilfunc
│       │   │   └── nilfunc.go
│       │   ├── print
│       │   │   └── print.go
│       │   ├── rangeloop
│       │   │   └── rangeloop.go
│       │   ├── shift
│       │   │   └── shift.go
│       │   ├── structtag
│       │   │   └── structtag.go
│       │   ├── tagtest
│       │   │   ├── file1.go
│       │   │   └── file2.go
│       │   ├── testingpkg
│       │   │   ├── tests.go
│       │   │   └── tests_test.go
│       │   ├── unmarshal
│       │   │   └── unmarshal.go
│       │   ├── unsafeptr
│       │   │   └── unsafeptr.go
│       │   └── unused
│       │       └── unused.go
│       └── vet_test.go
├── cmp.bash
├── compress
│   ├── bzip2
│   │   ├── bit_reader.go
│   │   ├── bzip2.go
│   │   ├── bzip2_test.go
│   │   ├── huffman.go
│   │   ├── move_to_front.go
│   │   └── testdata
│   │       ├── Isaac.Newton-Opticks.txt.bz2
│   │       ├── e.txt.bz2
│   │       ├── fail-issue5747.bz2
│   │       ├── pass-random1.bin
│   │       ├── pass-random1.bz2
│   │       ├── pass-random2.bin
│   │       ├── pass-random2.bz2
│   │       ├── pass-sawtooth.bz2
│   │       └── random.data.bz2
│   ├── flate
│   │   ├── deflate.go
│   │   ├── deflate_test.go
│   │   ├── deflatefast.go
│   │   ├── dict_decoder.go
│   │   ├── dict_decoder_test.go
│   │   ├── example_test.go
│   │   ├── flate_test.go
│   │   ├── huffman_bit_writer.go
│   │   ├── huffman_bit_writer_test.go
│   │   ├── huffman_code.go
│   │   ├── inflate.go
│   │   ├── inflate_test.go
│   │   ├── reader_test.go
│   │   ├── testdata
│   │   │   ├── huffman-null-max.dyn.expect
│   │   │   ├── huffman-null-max.dyn.expect-noinput
│   │   │   ├── huffman-null-max.golden
│   │   │   ├── huffman-null-max.in
│   │   │   ├── huffman-null-max.wb.expect
│   │   │   ├── huffman-null-max.wb.expect-noinput
│   │   │   ├── huffman-pi.dyn.expect
│   │   │   ├── huffman-pi.dyn.expect-noinput
│   │   │   ├── huffman-pi.golden
│   │   │   ├── huffman-pi.in
│   │   │   ├── huffman-pi.wb.expect
│   │   │   ├── huffman-pi.wb.expect-noinput
│   │   │   ├── huffman-rand-1k.dyn.expect
│   │   │   ├── huffman-rand-1k.dyn.expect-noinput
│   │   │   ├── huffman-rand-1k.golden
│   │   │   ├── huffman-rand-1k.in
│   │   │   ├── huffman-rand-1k.wb.expect
│   │   │   ├── huffman-rand-1k.wb.expect-noinput
│   │   │   ├── huffman-rand-limit.dyn.expect
│   │   │   ├── huffman-rand-limit.dyn.expect-noinput
│   │   │   ├── huffman-rand-limit.golden
│   │   │   ├── huffman-rand-limit.in
│   │   │   ├── huffman-rand-limit.wb.expect
│   │   │   ├── huffman-rand-limit.wb.expect-noinput
│   │   │   ├── huffman-rand-max.golden
│   │   │   ├── huffman-rand-max.in
│   │   │   ├── huffman-shifts.dyn.expect
│   │   │   ├── huffman-shifts.dyn.expect-noinput
│   │   │   ├── huffman-shifts.golden
│   │   │   ├── huffman-shifts.in
│   │   │   ├── huffman-shifts.wb.expect
│   │   │   ├── huffman-shifts.wb.expect-noinput
│   │   │   ├── huffman-text-shift.dyn.expect
│   │   │   ├── huffman-text-shift.dyn.expect-noinput
│   │   │   ├── huffman-text-shift.golden
│   │   │   ├── huffman-text-shift.in
│   │   │   ├── huffman-text-shift.wb.expect
│   │   │   ├── huffman-text-shift.wb.expect-noinput
│   │   │   ├── huffman-text.dyn.expect
│   │   │   ├── huffman-text.dyn.expect-noinput
│   │   │   ├── huffman-text.golden
│   │   │   ├── huffman-text.in
│   │   │   ├── huffman-text.wb.expect
│   │   │   ├── huffman-text.wb.expect-noinput
│   │   │   ├── huffman-zero.dyn.expect
│   │   │   ├── huffman-zero.dyn.expect-noinput
│   │   │   ├── huffman-zero.golden
│   │   │   ├── huffman-zero.in
│   │   │   ├── huffman-zero.wb.expect
│   │   │   ├── huffman-zero.wb.expect-noinput
│   │   │   ├── null-long-match.dyn.expect-noinput
│   │   │   └── null-long-match.wb.expect-noinput
│   │   ├── token.go
│   │   └── writer_test.go
│   ├── gzip
│   │   ├── example_test.go
│   │   ├── gunzip.go
│   │   ├── gunzip_test.go
│   │   ├── gzip.go
│   │   ├── gzip_test.go
│   │   ├── issue14937_test.go
│   │   └── testdata
│   │       └── issue6550.gz
│   ├── lzw
│   │   ├── reader.go
│   │   ├── reader_test.go
│   │   ├── writer.go
│   │   └── writer_test.go
│   ├── testdata
│   │   ├── e.txt
│   │   ├── gettysburg.txt
│   │   └── pi.txt
│   └── zlib
│       ├── example_test.go
│       ├── reader.go
│       ├── reader_test.go
│       ├── writer.go
│       └── writer_test.go
├── container
│   ├── heap
│   │   ├── example_intheap_test.go
│   │   ├── example_pq_test.go
│   │   ├── heap.go
│   │   └── heap_test.go
│   ├── list
│   │   ├── example_test.go
│   │   ├── list.go
│   │   └── list_test.go
│   └── ring
│       ├── example_test.go
│       ├── ring.go
│       └── ring_test.go
├── context
│   ├── benchmark_test.go
│   ├── context.go
│   ├── context_test.go
│   ├── example_test.go
│   ├── net_test.go
│   └── x_test.go
├── crypto
│   ├── aes
│   │   ├── aes_gcm.go
│   │   ├── aes_test.go
│   │   ├── asm_amd64.s
│   │   ├── asm_arm64.s
│   │   ├── asm_ppc64le.s
│   │   ├── asm_s390x.s
│   │   ├── block.go
│   │   ├── cbc_s390x.go
│   │   ├── cipher.go
│   │   ├── cipher_asm.go
│   │   ├── cipher_generic.go
│   │   ├── cipher_ppc64le.go
│   │   ├── cipher_s390x.go
│   │   ├── const.go
│   │   ├── ctr_s390x.go
│   │   ├── gcm_amd64.s
│   │   ├── gcm_arm64.s
│   │   ├── gcm_s390x.go
│   │   ├── modes.go
│   │   └── modes_test.go
│   ├── cipher
│   │   ├── benchmark_test.go
│   │   ├── cbc.go
│   │   ├── cbc_aes_test.go
│   │   ├── cfb.go
│   │   ├── cfb_test.go
│   │   ├── cipher.go
│   │   ├── cipher_test.go
│   │   ├── common_test.go
│   │   ├── ctr.go
│   │   ├── ctr_aes_test.go
│   │   ├── ctr_test.go
│   │   ├── example_test.go
│   │   ├── export_test.go
│   │   ├── gcm.go
│   │   ├── gcm_test.go
│   │   ├── io.go
│   │   ├── ofb.go
│   │   ├── ofb_test.go
│   │   ├── xor_amd64.go
│   │   ├── xor_amd64.s
│   │   ├── xor_generic.go
│   │   ├── xor_ppc64x.go
│   │   ├── xor_ppc64x.s
│   │   └── xor_test.go
│   ├── crypto.go
│   ├── des
│   │   ├── block.go
│   │   ├── cipher.go
│   │   ├── const.go
│   │   ├── des_test.go
│   │   └── example_test.go
│   ├── dsa
│   │   ├── dsa.go
│   │   └── dsa_test.go
│   ├── ecdsa
│   │   ├── ecdsa.go
│   │   ├── ecdsa_test.go
│   │   ├── example_test.go
│   │   └── testdata
│   │       └── SigVer.rsp.bz2
│   ├── ed25519
│   │   ├── ed25519.go
│   │   ├── ed25519_test.go
│   │   ├── internal
│   │   │   └── edwards25519
│   │   │       ├── const.go
│   │   │       └── edwards25519.go
│   │   └── testdata
│   │       └── sign.input.gz
│   ├── elliptic
│   │   ├── elliptic.go
│   │   ├── elliptic_test.go
│   │   ├── fuzz_test.go
│   │   ├── p224.go
│   │   ├── p224_test.go
│   │   ├── p256.go
│   │   ├── p256_asm.go
│   │   ├── p256_asm_amd64.s
│   │   ├── p256_asm_arm64.s
│   │   ├── p256_asm_s390x.s
│   │   ├── p256_generic.go
│   │   └── p256_s390x.go
│   ├── hmac
│   │   ├── hmac.go
│   │   └── hmac_test.go
│   ├── internal
│   │   ├── randutil
│   │   │   └── randutil.go
│   │   └── subtle
│   │       ├── aliasing.go
│   │       ├── aliasing_appengine.go
│   │       └── aliasing_test.go
│   ├── issue21104_test.go
│   ├── md5
│   │   ├── example_test.go
│   │   ├── gen.go
│   │   ├── md5.go
│   │   ├── md5_test.go
│   │   ├── md5block.go
│   │   ├── md5block_386.s
│   │   ├── md5block_amd64.s
│   │   ├── md5block_amd64p32.s
│   │   ├── md5block_arm.s
│   │   ├── md5block_arm64.s
│   │   ├── md5block_decl.go
│   │   ├── md5block_generic.go
│   │   ├── md5block_ppc64x.s
│   │   └── md5block_s390x.s
│   ├── rand
│   │   ├── eagain.go
│   │   ├── example_test.go
│   │   ├── rand.go
│   │   ├── rand_batched.go
│   │   ├── rand_batched_test.go
│   │   ├── rand_freebsd.go
│   │   ├── rand_js.go
│   │   ├── rand_linux.go
│   │   ├── rand_openbsd.go
│   │   ├── rand_test.go
│   │   ├── rand_unix.go
│   │   ├── rand_windows.go
│   │   ├── util.go
│   │   └── util_test.go
│   ├── rc4
│   │   ├── rc4.go
│   │   └── rc4_test.go
│   ├── rsa
│   │   ├── example_test.go
│   │   ├── pkcs1v15.go
│   │   ├── pkcs1v15_test.go
│   │   ├── pss.go
│   │   ├── pss_test.go
│   │   ├── rsa.go
│   │   ├── rsa_test.go
│   │   └── testdata
│   │       └── pss-vect.txt.bz2
│   ├── sha1
│   │   ├── example_test.go
│   │   ├── fallback_test.go
│   │   ├── issue15617_test.go
│   │   ├── sha1.go
│   │   ├── sha1_test.go
│   │   ├── sha1block.go
│   │   ├── sha1block_386.s
│   │   ├── sha1block_amd64.go
│   │   ├── sha1block_amd64.s
│   │   ├── sha1block_amd64p32.s
│   │   ├── sha1block_arm.s
│   │   ├── sha1block_arm64.go
│   │   ├── sha1block_arm64.s
│   │   ├── sha1block_decl.go
│   │   ├── sha1block_generic.go
│   │   ├── sha1block_s390x.go
│   │   └── sha1block_s390x.s
│   ├── sha256
│   │   ├── example_test.go
│   │   ├── fallback_test.go
│   │   ├── sha256.go
│   │   ├── sha256_test.go
│   │   ├── sha256block.go
│   │   ├── sha256block_386.s
│   │   ├── sha256block_amd64.go
│   │   ├── sha256block_amd64.s
│   │   ├── sha256block_arm64.go
│   │   ├── sha256block_arm64.s
│   │   ├── sha256block_decl.go
│   │   ├── sha256block_generic.go
│   │   ├── sha256block_ppc64le.s
│   │   ├── sha256block_s390x.go
│   │   └── sha256block_s390x.s
│   ├── sha512
│   │   ├── fallback_test.go
│   │   ├── sha512.go
│   │   ├── sha512_test.go
│   │   ├── sha512block.go
│   │   ├── sha512block_amd64.go
│   │   ├── sha512block_amd64.s
│   │   ├── sha512block_decl.go
│   │   ├── sha512block_generic.go
│   │   ├── sha512block_ppc64le.s
│   │   ├── sha512block_s390x.go
│   │   └── sha512block_s390x.s
│   ├── subtle
│   │   ├── constant_time.go
│   │   └── constant_time_test.go
│   ├── tls
│   │   ├── alert.go
│   │   ├── auth.go
│   │   ├── auth_test.go
│   │   ├── cipher_suites.go
│   │   ├── common.go
│   │   ├── conn.go
│   │   ├── conn_test.go
│   │   ├── example_test.go
│   │   ├── generate_cert.go
│   │   ├── handshake_client.go
│   │   ├── handshake_client_test.go
│   │   ├── handshake_client_tls13.go
│   │   ├── handshake_messages.go
│   │   ├── handshake_messages_test.go
│   │   ├── handshake_server.go
│   │   ├── handshake_server_test.go
│   │   ├── handshake_server_tls13.go
│   │   ├── handshake_test.go
│   │   ├── key_agreement.go
│   │   ├── key_schedule.go
│   │   ├── key_schedule_test.go
│   │   ├── prf.go
│   │   ├── prf_test.go
│   │   ├── testdata
│   │   │   ├── Client-TLSv10-ClientCert-ECDSA-ECDSA
│   │   │   ├── Client-TLSv10-ClientCert-ECDSA-RSA
│   │   │   ├── Client-TLSv10-ClientCert-Ed25519
│   │   │   ├── Client-TLSv10-ClientCert-RSA-ECDSA
│   │   │   ├── Client-TLSv10-ClientCert-RSA-RSA
│   │   │   ├── Client-TLSv10-ECDHE-ECDSA-AES
│   │   │   ├── Client-TLSv10-ECDHE-RSA-AES
│   │   │   ├── Client-TLSv10-Ed25519
│   │   │   ├── Client-TLSv10-ExportKeyingMaterial
│   │   │   ├── Client-TLSv10-RSA-RC4
│   │   │   ├── Client-TLSv11-ECDHE-ECDSA-AES
│   │   │   ├── Client-TLSv11-ECDHE-RSA-AES
│   │   │   ├── Client-TLSv11-Ed25519
│   │   │   ├── Client-TLSv11-RSA-RC4
│   │   │   ├── Client-TLSv12-AES128-GCM-SHA256
│   │   │   ├── Client-TLSv12-AES128-SHA256
│   │   │   ├── Client-TLSv12-AES256-GCM-SHA384
│   │   │   ├── Client-TLSv12-ALPN
│   │   │   ├── Client-TLSv12-ALPN-NoMatch
│   │   │   ├── Client-TLSv12-ClientCert-ECDSA-ECDSA
│   │   │   ├── Client-TLSv12-ClientCert-ECDSA-RSA
│   │   │   ├── Client-TLSv12-ClientCert-Ed25519
│   │   │   ├── Client-TLSv12-ClientCert-RSA-AES256-GCM-SHA384
│   │   │   ├── Client-TLSv12-ClientCert-RSA-ECDSA
│   │   │   ├── Client-TLSv12-ClientCert-RSA-RSA
│   │   │   ├── Client-TLSv12-ClientCert-RSA-RSAPKCS1v15
│   │   │   ├── Client-TLSv12-ClientCert-RSA-RSAPSS
│   │   │   ├── Client-TLSv12-ECDHE-ECDSA-AES
│   │   │   ├── Client-TLSv12-ECDHE-ECDSA-AES-GCM
│   │   │   ├── Client-TLSv12-ECDHE-ECDSA-AES128-SHA256
│   │   │   ├── Client-TLSv12-ECDHE-ECDSA-AES256-GCM-SHA384
│   │   │   ├── Client-TLSv12-ECDHE-ECDSA-CHACHA20-POLY1305
│   │   │   ├── Client-TLSv12-ECDHE-RSA-AES
│   │   │   ├── Client-TLSv12-ECDHE-RSA-AES128-SHA256
│   │   │   ├── Client-TLSv12-ECDHE-RSA-CHACHA20-POLY1305
│   │   │   ├── Client-TLSv12-Ed25519
│   │   │   ├── Client-TLSv12-ExportKeyingMaterial
│   │   │   ├── Client-TLSv12-P256-ECDHE
│   │   │   ├── Client-TLSv12-RSA-RC4
│   │   │   ├── Client-TLSv12-RenegotiateOnce
│   │   │   ├── Client-TLSv12-RenegotiateTwice
│   │   │   ├── Client-TLSv12-RenegotiateTwiceRejected
│   │   │   ├── Client-TLSv12-RenegotiationRejected
│   │   │   ├── Client-TLSv12-SCT
│   │   │   ├── Client-TLSv12-X25519-ECDHE
│   │   │   ├── Client-TLSv13-AES128-SHA256
│   │   │   ├── Client-TLSv13-AES256-SHA384
│   │   │   ├── Client-TLSv13-ALPN
│   │   │   ├── Client-TLSv13-CHACHA20-SHA256
│   │   │   ├── Client-TLSv13-ClientCert-ECDSA-RSA
│   │   │   ├── Client-TLSv13-ClientCert-Ed25519
│   │   │   ├── Client-TLSv13-ClientCert-RSA-ECDSA
│   │   │   ├── Client-TLSv13-ClientCert-RSA-RSAPSS
│   │   │   ├── Client-TLSv13-ECDSA
│   │   │   ├── Client-TLSv13-Ed25519
│   │   │   ├── Client-TLSv13-ExportKeyingMaterial
│   │   │   ├── Client-TLSv13-HelloRetryRequest
│   │   │   ├── Client-TLSv13-KeyUpdate
│   │   │   ├── Client-TLSv13-P256-ECDHE
│   │   │   ├── Client-TLSv13-X25519-ECDHE
│   │   │   ├── Server-SSLv3-RSA-3DES
│   │   │   ├── Server-SSLv3-RSA-AES
│   │   │   ├── Server-SSLv3-RSA-RC4
│   │   │   ├── Server-TLSv10-ECDHE-ECDSA-AES
│   │   │   ├── Server-TLSv10-ExportKeyingMaterial
│   │   │   ├── Server-TLSv10-RSA-3DES
│   │   │   ├── Server-TLSv10-RSA-AES
│   │   │   ├── Server-TLSv10-RSA-RC4
│   │   │   ├── Server-TLSv11-FallbackSCSV
│   │   │   ├── Server-TLSv11-RSA-RC4
│   │   │   ├── Server-TLSv12-ALPN
│   │   │   ├── Server-TLSv12-ALPN-NoMatch
│   │   │   ├── Server-TLSv12-CipherSuiteCertPreferenceECDSA
│   │   │   ├── Server-TLSv12-CipherSuiteCertPreferenceRSA
│   │   │   ├── Server-TLSv12-ClientAuthRequestedAndECDSAGiven
│   │   │   ├── Server-TLSv12-ClientAuthRequestedAndEd25519Given
│   │   │   ├── Server-TLSv12-ClientAuthRequestedAndGiven
│   │   │   ├── Server-TLSv12-ClientAuthRequestedAndPKCS1v15Given
│   │   │   ├── Server-TLSv12-ClientAuthRequestedNotGiven
│   │   │   ├── Server-TLSv12-ECDHE-ECDSA-AES
│   │   │   ├── Server-TLSv12-Ed25519
│   │   │   ├── Server-TLSv12-ExportKeyingMaterial
│   │   │   ├── Server-TLSv12-IssueTicket
│   │   │   ├── Server-TLSv12-IssueTicketPreDisable
│   │   │   ├── Server-TLSv12-P256
│   │   │   ├── Server-TLSv12-RSA-3DES
│   │   │   ├── Server-TLSv12-RSA-AES
│   │   │   ├── Server-TLSv12-RSA-AES-GCM
│   │   │   ├── Server-TLSv12-RSA-AES256-GCM-SHA384
│   │   │   ├── Server-TLSv12-RSA-RC4
│   │   │   ├── Server-TLSv12-RSA-RSAPKCS1v15
│   │   │   ├── Server-TLSv12-RSA-RSAPSS
│   │   │   ├── Server-TLSv12-Resume
│   │   │   ├── Server-TLSv12-ResumeDisabled
│   │   │   ├── Server-TLSv12-SNI
│   │   │   ├── Server-TLSv12-SNI-GetCertificate
│   │   │   ├── Server-TLSv12-SNI-GetCertificateNotFound
│   │   │   ├── Server-TLSv12-X25519
│   │   │   ├── Server-TLSv13-AES128-SHA256
│   │   │   ├── Server-TLSv13-AES256-SHA384
│   │   │   ├── Server-TLSv13-ALPN
│   │   │   ├── Server-TLSv13-ALPN-NoMatch
│   │   │   ├── Server-TLSv13-CHACHA20-SHA256
│   │   │   ├── Server-TLSv13-ClientAuthRequestedAndECDSAGiven
│   │   │   ├── Server-TLSv13-ClientAuthRequestedAndEd25519Given
│   │   │   ├── Server-TLSv13-ClientAuthRequestedAndGiven
│   │   │   ├── Server-TLSv13-ClientAuthRequestedNotGiven
│   │   │   ├── Server-TLSv13-ECDHE-ECDSA-AES
│   │   │   ├── Server-TLSv13-Ed25519
│   │   │   ├── Server-TLSv13-ExportKeyingMaterial
│   │   │   ├── Server-TLSv13-HelloRetryRequest
│   │   │   ├── Server-TLSv13-IssueTicket
│   │   │   ├── Server-TLSv13-IssueTicketPreDisable
│   │   │   ├── Server-TLSv13-P256
│   │   │   ├── Server-TLSv13-RSA-RSAPSS
│   │   │   ├── Server-TLSv13-Resume
│   │   │   ├── Server-TLSv13-Resume-HelloRetryRequest
│   │   │   ├── Server-TLSv13-ResumeDisabled
│   │   │   ├── Server-TLSv13-X25519
│   │   │   ├── example-cert.pem
│   │   │   └── example-key.pem
│   │   ├── ticket.go
│   │   ├── tls.go
│   │   └── tls_test.go
│   └── x509
│       ├── cert_pool.go
│       ├── example_test.go
│       ├── name_constraints_test.go
│       ├── pem_decrypt.go
│       ├── pem_decrypt_test.go
│       ├── pkcs1.go
│       ├── pkcs8.go
│       ├── pkcs8_test.go
│       ├── pkix
│       │   └── pkix.go
│       ├── root.go
│       ├── root_aix.go
│       ├── root_bsd.go
│       ├── root_cgo_darwin.go
│       ├── root_darwin.go
│       ├── root_darwin_arm_gen.go
│       ├── root_darwin_armx.go
│       ├── root_darwin_test.go
│       ├── root_js.go
│       ├── root_linux.go
│       ├── root_nacl.go
│       ├── root_nocgo_darwin.go
│       ├── root_plan9.go
│       ├── root_solaris.go
│       ├── root_unix.go
│       ├── root_unix_test.go
│       ├── root_windows.go
│       ├── sec1.go
│       ├── sec1_test.go
│       ├── test-file.crt
│       ├── testdata
│       │   └── test-dir.crt
│       ├── verify.go
│       ├── verify_test.go
│       ├── x509.go
│       ├── x509_test.go
│       └── x509_test_import.go
├── database
│   └── sql
│       ├── convert.go
│       ├── convert_test.go
│       ├── ctxutil.go
│       ├── doc.txt
│       ├── driver
│       │   ├── driver.go
│       │   ├── types.go
│       │   └── types_test.go
│       ├── example_cli_test.go
│       ├── example_service_test.go
│       ├── example_test.go
│       ├── fakedb_test.go
│       ├── sql.go
│       └── sql_test.go
├── debug
│   ├── dwarf
│   │   ├── attr_string.go
│   │   ├── buf.go
│   │   ├── class_string.go
│   │   ├── const.go
│   │   ├── entry.go
│   │   ├── entry_test.go
│   │   ├── export_test.go
│   │   ├── line.go
│   │   ├── line_test.go
│   │   ├── open.go
│   │   ├── tag_string.go
│   │   ├── testdata
│   │   │   ├── cppunsuptypes.cc
│   │   │   ├── cppunsuptypes.elf
│   │   │   ├── cycle.c
│   │   │   ├── cycle.elf
│   │   │   ├── line-clang.elf
│   │   │   ├── line-gcc-win.bin
│   │   │   ├── line-gcc.elf
│   │   │   ├── line1.c
│   │   │   ├── line1.h
│   │   │   ├── line2.c
│   │   │   ├── ranges.c
│   │   │   ├── ranges.elf
│   │   │   ├── split.c
│   │   │   ├── split.elf
│   │   │   ├── typedef.c
│   │   │   ├── typedef.elf
│   │   │   ├── typedef.elf4
│   │   │   └── typedef.macho
│   │   ├── type.go
│   │   ├── type_test.go
│   │   ├── typeunit.go
│   │   └── unit.go
│   ├── elf
│   │   ├── elf.go
│   │   ├── elf_test.go
│   │   ├── file.go
│   │   ├── file_test.go
│   │   ├── reader.go
│   │   ├── symbols_test.go
│   │   └── testdata
│   │       ├── compressed-32.obj
│   │       ├── compressed-64.obj
│   │       ├── gcc-386-freebsd-exec
│   │       ├── gcc-amd64-linux-exec
│   │       ├── gcc-amd64-openbsd-debug-with-rela.obj
│   │       ├── go-relocation-test-clang-arm.obj
│   │       ├── go-relocation-test-clang-x86.obj
│   │       ├── go-relocation-test-gcc424-x86-64.obj
│   │       ├── go-relocation-test-gcc441-x86-64.obj
│   │       ├── go-relocation-test-gcc441-x86.obj
│   │       ├── go-relocation-test-gcc482-aarch64.obj
│   │       ├── go-relocation-test-gcc482-ppc64le.obj
│   │       ├── go-relocation-test-gcc492-arm.obj
│   │       ├── go-relocation-test-gcc492-mips64.obj
│   │       ├── go-relocation-test-gcc492-mipsle.obj
│   │       ├── go-relocation-test-gcc493-mips64le.obj
│   │       ├── go-relocation-test-gcc5-ppc.obj
│   │       ├── go-relocation-test-gcc531-s390x.obj
│   │       ├── go-relocation-test-gcc540-mips.obj
│   │       ├── go-relocation-test-gcc620-sparc64.obj
│   │       ├── go-relocation-test-gcc720-riscv64.obj
│   │       ├── hello-world-core.gz
│   │       ├── hello.c
│   │       └── zdebug-test-gcc484-x86-64.obj
│   ├── gosym
│   │   ├── pclntab.go
│   │   ├── pclntab_test.go
│   │   ├── symtab.go
│   │   ├── symtab_test.go
│   │   └── testdata
│   │       ├── main.go
│   │       ├── pclinetest.h
│   │       └── pclinetest.s
│   ├── macho
│   │   ├── fat.go
│   │   ├── file.go
│   │   ├── file_test.go
│   │   ├── macho.go
│   │   ├── reloctype.go
│   │   ├── reloctype_string.go
│   │   └── testdata
│   │       ├── clang-386-darwin-exec-with-rpath
│   │       ├── clang-386-darwin.obj
│   │       ├── clang-amd64-darwin-exec-with-rpath
│   │       ├── clang-amd64-darwin.obj
│   │       ├── fat-gcc-386-amd64-darwin-exec
│   │       ├── gcc-386-darwin-exec
│   │       ├── gcc-amd64-darwin-exec
│   │       ├── gcc-amd64-darwin-exec-debug
│   │       └── hello.c
│   ├── pe
│   │   ├── file.go
│   │   ├── file_cgo_test.go
│   │   ├── file_test.go
│   │   ├── pe.go
│   │   ├── section.go
│   │   ├── string.go
│   │   ├── symbol.go
│   │   └── testdata
│   │       ├── gcc-386-mingw-exec
│   │       ├── gcc-386-mingw-no-symbols-exec
│   │       ├── gcc-386-mingw-obj
│   │       ├── gcc-amd64-mingw-exec
│   │       ├── gcc-amd64-mingw-obj
│   │       └── hello.c
│   └── plan9obj
│       ├── file.go
│       ├── file_test.go
│       ├── plan9obj.go
│       └── testdata
│           ├── 386-plan9-exec
│           ├── amd64-plan9-exec
│           └── hello.c
├── encoding
│   ├── ascii85
│   │   ├── ascii85.go
│   │   └── ascii85_test.go
│   ├── asn1
│   │   ├── asn1.go
│   │   ├── asn1_test.go
│   │   ├── common.go
│   │   ├── marshal.go
│   │   └── marshal_test.go
│   ├── base32
│   │   ├── base32.go
│   │   ├── base32_test.go
│   │   └── example_test.go
│   ├── base64
│   │   ├── base64.go
│   │   ├── base64_test.go
│   │   └── example_test.go
│   ├── binary
│   │   ├── binary.go
│   │   ├── binary_test.go
│   │   ├── example_test.go
│   │   ├── varint.go
│   │   └── varint_test.go
│   ├── csv
│   │   ├── example_test.go
│   │   ├── fuzz.go
│   │   ├── reader.go
│   │   ├── reader_test.go
│   │   ├── writer.go
│   │   └── writer_test.go
│   ├── encoding.go
│   ├── gob
│   │   ├── codec_test.go
│   │   ├── debug.go
│   │   ├── dec_helpers.go
│   │   ├── decgen.go
│   │   ├── decode.go
│   │   ├── decoder.go
│   │   ├── doc.go
│   │   ├── dump.go
│   │   ├── enc_helpers.go
│   │   ├── encgen.go
│   │   ├── encode.go
│   │   ├── encoder.go
│   │   ├── encoder_test.go
│   │   ├── error.go
│   │   ├── example_encdec_test.go
│   │   ├── example_interface_test.go
│   │   ├── example_test.go
│   │   ├── gobencdec_test.go
│   │   ├── timing_test.go
│   │   ├── type.go
│   │   └── type_test.go
│   ├── hex
│   │   ├── example_test.go
│   │   ├── hex.go
│   │   └── hex_test.go
│   ├── json
│   │   ├── bench_test.go
│   │   ├── decode.go
│   │   ├── decode_test.go
│   │   ├── encode.go
│   │   ├── encode_test.go
│   │   ├── example_marshaling_test.go
│   │   ├── example_test.go
│   │   ├── example_text_marshaling_test.go
│   │   ├── fold.go
│   │   ├── fold_test.go
│   │   ├── fuzz.go
│   │   ├── indent.go
│   │   ├── number_test.go
│   │   ├── scanner.go
│   │   ├── scanner_test.go
│   │   ├── stream.go
│   │   ├── stream_test.go
│   │   ├── tables.go
│   │   ├── tagkey_test.go
│   │   ├── tags.go
│   │   ├── tags_test.go
│   │   └── testdata
│   │       └── code.json.gz
│   ├── pem
│   │   ├── example_test.go
│   │   ├── pem.go
│   │   └── pem_test.go
│   └── xml
│       ├── atom_test.go
│       ├── example_marshaling_test.go
│       ├── example_test.go
│       ├── example_text_marshaling_test.go
│       ├── marshal.go
│       ├── marshal_test.go
│       ├── read.go
│       ├── read_test.go
│       ├── typeinfo.go
│       ├── xml.go
│       └── xml_test.go
├── errors
│   ├── errors.go
│   ├── errors_test.go
│   ├── example_test.go
│   ├── wrap.go
│   └── wrap_test.go
├── expvar
│   ├── expvar.go
│   └── expvar_test.go
├── flag
│   ├── example_test.go
│   ├── example_value_test.go
│   ├── export_test.go
│   ├── flag.go
│   └── flag_test.go
├── fmt
│   ├── doc.go
│   ├── errors.go
│   ├── errors_test.go
│   ├── example_test.go
│   ├── export_test.go
│   ├── fmt_test.go
│   ├── format.go
│   ├── gostringer_example_test.go
│   ├── print.go
│   ├── scan.go
│   ├── scan_test.go
│   ├── stringer_example_test.go
│   └── stringer_test.go
├── go
│   ├── ast
│   │   ├── ast.go
│   │   ├── ast_test.go
│   │   ├── commentmap.go
│   │   ├── commentmap_test.go
│   │   ├── example_test.go
│   │   ├── filter.go
│   │   ├── filter_test.go
│   │   ├── import.go
│   │   ├── print.go
│   │   ├── print_test.go
│   │   ├── resolve.go
│   │   ├── scope.go
│   │   └── walk.go
│   ├── build
│   │   ├── build.go
│   │   ├── build_test.go
│   │   ├── deps_test.go
│   │   ├── doc.go
│   │   ├── gc.go
│   │   ├── gccgo.go
│   │   ├── read.go
│   │   ├── read_test.go
│   │   ├── syslist.go
│   │   ├── syslist_test.go
│   │   ├── testdata
│   │   │   ├── doc
│   │   │   │   ├── a_test.go
│   │   │   │   ├── b_test.go
│   │   │   │   ├── c_test.go
│   │   │   │   ├── d_test.go
│   │   │   │   ├── e.go
│   │   │   │   └── f.go
│   │   │   ├── empty
│   │   │   │   └── dummy
│   │   │   ├── multi
│   │   │   │   ├── file.go
│   │   │   │   └── file_appengine.go
│   │   │   ├── other
│   │   │   │   ├── file
│   │   │   │   │   └── file.go
│   │   │   │   └── main.go
│   │   │   └── withvendor
│   │   │       └── src
│   │   │           └── a
│   │   │               ├── b
│   │   │               │   └── b.go
│   │   │               └── vendor
│   │   │                   └── c
│   │   │                       └── d
│   │   │                           └── d.go
│   │   └── zcgo.go
│   ├── constant
│   │   ├── example_test.go
│   │   ├── value.go
│   │   └── value_test.go
│   ├── doc
│   │   ├── Makefile
│   │   ├── comment.go
│   │   ├── comment_test.go
│   │   ├── doc.go
│   │   ├── doc_test.go
│   │   ├── example.go
│   │   ├── example_test.go
│   │   ├── exports.go
│   │   ├── filter.go
│   │   ├── headscan.go
│   │   ├── reader.go
│   │   ├── synopsis.go
│   │   ├── synopsis_test.go
│   │   └── testdata
│   │       ├── a.0.golden
│   │       ├── a.1.golden
│   │       ├── a.2.golden
│   │       ├── a0.go
│   │       ├── a1.go
│   │       ├── b.0.golden
│   │       ├── b.1.golden
│   │       ├── b.2.golden
│   │       ├── b.go
│   │       ├── benchmark.go
│   │       ├── blank.0.golden
│   │       ├── blank.1.golden
│   │       ├── blank.2.golden
│   │       ├── blank.go
│   │       ├── bugpara.0.golden
│   │       ├── bugpara.1.golden
│   │       ├── bugpara.2.golden
│   │       ├── bugpara.go
│   │       ├── c.0.golden
│   │       ├── c.1.golden
│   │       ├── c.2.golden
│   │       ├── c.go
│   │       ├── d.0.golden
│   │       ├── d.1.golden
│   │       ├── d.2.golden
│   │       ├── d1.go
│   │       ├── d2.go
│   │       ├── e.0.golden
│   │       ├── e.1.golden
│   │       ├── e.2.golden
│   │       ├── e.go
│   │       ├── error1.0.golden
│   │       ├── error1.1.golden
│   │       ├── error1.2.golden
│   │       ├── error1.go
│   │       ├── error2.0.golden
│   │       ├── error2.1.golden
│   │       ├── error2.2.golden
│   │       ├── error2.go
│   │       ├── example.go
│   │       ├── f.0.golden
│   │       ├── f.1.golden
│   │       ├── f.2.golden
│   │       ├── f.go
│   │       ├── g.0.golden
│   │       ├── g.1.golden
│   │       ├── g.2.golden
│   │       ├── g.go
│   │       ├── issue12839.0.golden
│   │       ├── issue12839.1.golden
│   │       ├── issue12839.2.golden
│   │       ├── issue12839.go
│   │       ├── issue13742.0.golden
│   │       ├── issue13742.1.golden
│   │       ├── issue13742.2.golden
│   │       ├── issue13742.go
│   │       ├── issue16153.0.golden
│   │       ├── issue16153.1.golden
│   │       ├── issue16153.2.golden
│   │       ├── issue16153.go
│   │       ├── issue17788.0.golden
│   │       ├── issue17788.1.golden
│   │       ├── issue17788.2.golden
│   │       ├── issue17788.go
│   │       ├── issue22856.0.golden
│   │       ├── issue22856.1.golden
│   │       ├── issue22856.2.golden
│   │       ├── issue22856.go
│   │       ├── predeclared.0.golden
│   │       ├── predeclared.1.golden
│   │       ├── predeclared.2.golden
│   │       ├── predeclared.go
│   │       ├── template.txt
│   │       ├── testing.0.golden
│   │       ├── testing.1.golden
│   │       ├── testing.2.golden
│   │       └── testing.go
│   ├── format
│   │   ├── benchmark_test.go
│   │   ├── example_test.go
│   │   ├── format.go
│   │   ├── format_test.go
│   │   └── internal.go
│   ├── importer
│   │   ├── importer.go
│   │   └── importer_test.go
│   ├── internal
│   │   ├── gccgoimporter
│   │   │   ├── ar.go
│   │   │   ├── gccgoinstallation.go
│   │   │   ├── gccgoinstallation_test.go
│   │   │   ├── importer.go
│   │   │   ├── importer_test.go
│   │   │   ├── parser.go
│   │   │   ├── parser_test.go
│   │   │   └── testdata
│   │   │       ├── aliases.go
│   │   │       ├── aliases.gox
│   │   │       ├── complexnums.go
│   │   │       ├── complexnums.gox
│   │   │       ├── conversions.go
│   │   │       ├── conversions.gox
│   │   │       ├── escapeinfo.go
│   │   │       ├── escapeinfo.gox
│   │   │       ├── imports.go
│   │   │       ├── imports.gox
│   │   │       ├── issue27856.go
│   │   │       ├── issue27856.gox
│   │   │       ├── issue29198.go
│   │   │       ├── issue29198.gox
│   │   │       ├── issue30628.go
│   │   │       ├── issue30628.gox
│   │   │       ├── issue31540.go
│   │   │       ├── issue31540.gox
│   │   │       ├── libimportsar.a
│   │   │       ├── nointerface.go
│   │   │       ├── nointerface.gox
│   │   │       ├── pointer.go
│   │   │       ├── pointer.gox
│   │   │       ├── time.gox
│   │   │       ├── unicode.gox
│   │   │       └── v1reflect.gox
│   │   ├── gcimporter
│   │   │   ├── bimport.go
│   │   │   ├── exportdata.go
│   │   │   ├── gcimporter.go
│   │   │   ├── gcimporter_test.go
│   │   │   ├── iimport.go
│   │   │   └── testdata
│   │   │       ├── a.go
│   │   │       ├── b.go
│   │   │       ├── exports.go
│   │   │       ├── issue15920.go
│   │   │       ├── issue20046.go
│   │   │       ├── issue25301.go
│   │   │       ├── issue25596.go
│   │   │       ├── p.go
│   │   │       └── versions
│   │   │           ├── test.go
│   │   │           ├── test_go1.11_0i.a
│   │   │           ├── test_go1.11_6b.a
│   │   │           ├── test_go1.11_999b.a
│   │   │           ├── test_go1.11_999i.a
│   │   │           ├── test_go1.7_0.a
│   │   │           ├── test_go1.7_1.a
│   │   │           ├── test_go1.8_4.a
│   │   │           └── test_go1.8_5.a
│   │   └── srcimporter
│   │       ├── srcimporter.go
│   │       ├── srcimporter_test.go
│   │       └── testdata
│   │           ├── issue20855
│   │           │   └── issue20855.go
│   │           ├── issue23092
│   │           │   └── issue23092.go
│   │           └── issue24392
│   │               └── issue24392.go
│   ├── parser
│   │   ├── error_test.go
│   │   ├── example_test.go
│   │   ├── interface.go
│   │   ├── parser.go
│   │   ├── parser_test.go
│   │   ├── performance_test.go
│   │   ├── short_test.go
│   │   └── testdata
│   │       ├── commas.src
│   │       ├── issue11377.src
│   │       ├── issue23434.src
│   │       └── issue3106.src
│   ├── printer
│   │   ├── example_test.go
│   │   ├── nodes.go
│   │   ├── performance_test.go
│   │   ├── printer.go
│   │   ├── printer_test.go
│   │   └── testdata
│   │       ├── alignment.golden
│   │       ├── alignment.input
│   │       ├── comments.golden
│   │       ├── comments.input
│   │       ├── comments.x
│   │       ├── comments2.golden
│   │       ├── comments2.input
│   │       ├── complit.input
│   │       ├── complit.x
│   │       ├── declarations.golden
│   │       ├── declarations.input
│   │       ├── empty.golden
│   │       ├── empty.input
│   │       ├── expressions.golden
│   │       ├── expressions.input
│   │       ├── expressions.raw
│   │       ├── linebreaks.golden
│   │       ├── linebreaks.input
│   │       ├── parser.go
│   │       ├── slow.golden
│   │       ├── slow.input
│   │       ├── statements.golden
│   │       └── statements.input
│   ├── scanner
│   │   ├── errors.go
│   │   ├── example_test.go
│   │   ├── scanner.go
│   │   └── scanner_test.go
│   ├── token
│   │   ├── example_test.go
│   │   ├── position.go
│   │   ├── position_test.go
│   │   ├── serialize.go
│   │   ├── serialize_test.go
│   │   ├── token.go
│   │   └── token_test.go
│   └── types
│       ├── api.go
│       ├── api_test.go
│       ├── assignments.go
│       ├── builtins.go
│       ├── builtins_test.go
│       ├── call.go
│       ├── check.go
│       ├── check_test.go
│       ├── conversions.go
│       ├── decl.go
│       ├── errors.go
│       ├── eval.go
│       ├── eval_test.go
│       ├── example_test.go
│       ├── expr.go
│       ├── exprstring.go
│       ├── exprstring_test.go
│       ├── gccgosizes.go
│       ├── gotype.go
│       ├── hilbert_test.go
│       ├── initorder.go
│       ├── interfaces.go
│       ├── issues_test.go
│       ├── labels.go
│       ├── lookup.go
│       ├── methodset.go
│       ├── object.go
│       ├── object_test.go
│       ├── objset.go
│       ├── operand.go
│       ├── package.go
│       ├── predicates.go
│       ├── resolver.go
│       ├── resolver_test.go
│       ├── return.go
│       ├── scope.go
│       ├── selection.go
│       ├── self_test.go
│       ├── sizes.go
│       ├── sizes_test.go
│       ├── stdlib_test.go
│       ├── stmt.go
│       ├── testdata
│       │   ├── blank.src
│       │   ├── builtins.src
│       │   ├── const0.src
│       │   ├── const1.src
│       │   ├── constdecl.src
│       │   ├── conversions.src
│       │   ├── conversions2.src
│       │   ├── cycles.src
│       │   ├── cycles1.src
│       │   ├── cycles2.src
│       │   ├── cycles3.src
│       │   ├── cycles4.src
│       │   ├── cycles5.src
│       │   ├── decls0.src
│       │   ├── decls1.src
│       │   ├── decls2a.src
│       │   ├── decls2b.src
│       │   ├── decls3.src
│       │   ├── decls4.src
│       │   ├── decls5.src
│       │   ├── errors.src
│       │   ├── expr0.src
│       │   ├── expr1.src
│       │   ├── expr2.src
│       │   ├── expr3.src
│       │   ├── gotos.src
│       │   ├── importC.src
│       │   ├── importdecl0a.src
│       │   ├── importdecl0b.src
│       │   ├── importdecl1a.src
│       │   ├── importdecl1b.src
│       │   ├── init0.src
│       │   ├── init1.src
│       │   ├── init2.src
│       │   ├── issue23203a.src
│       │   ├── issue23203b.src
│       │   ├── issue25008a.src
│       │   ├── issue25008b.src
│       │   ├── issue26390.src
│       │   ├── issue28251.src
│       │   ├── issues.src
│       │   ├── labels.src
│       │   ├── literals.src
│       │   ├── methodsets.src
│       │   ├── shifts.src
│       │   ├── stmt0.src
│       │   ├── stmt1.src
│       │   └── vardecl.src
│       ├── token_test.go
│       ├── type.go
│       ├── typestring.go
│       ├── typestring_test.go
│       ├── typexpr.go
│       └── universe.go
├── go.mod
├── go.sum
├── hash
│   ├── adler32
│   │   ├── adler32.go
│   │   └── adler32_test.go
│   ├── crc32
│   │   ├── crc32.go
│   │   ├── crc32_amd64.go
│   │   ├── crc32_amd64.s
│   │   ├── crc32_amd64p32.go
│   │   ├── crc32_amd64p32.s
│   │   ├── crc32_arm64.go
│   │   ├── crc32_arm64.s
│   │   ├── crc32_generic.go
│   │   ├── crc32_otherarch.go
│   │   ├── crc32_ppc64le.go
│   │   ├── crc32_ppc64le.s
│   │   ├── crc32_s390x.go
│   │   ├── crc32_s390x.s
│   │   ├── crc32_table_ppc64le.s
│   │   ├── crc32_test.go
│   │   ├── example_test.go
│   │   └── gen_const_ppc64le.go
│   ├── crc64
│   │   ├── crc64.go
│   │   └── crc64_test.go
│   ├── example_test.go
│   ├── fnv
│   │   ├── fnv.go
│   │   └── fnv_test.go
│   ├── hash.go
│   ├── marshal_test.go
│   ├── test_cases.txt
│   └── test_gen.awk
├── html
│   ├── entity.go
│   ├── entity_test.go
│   ├── escape.go
│   ├── escape_test.go
│   ├── example_test.go
│   ├── fuzz.go
│   └── template
│       ├── attr.go
│       ├── attr_string.go
│       ├── clone_test.go
│       ├── content.go
│       ├── content_test.go
│       ├── context.go
│       ├── css.go
│       ├── css_test.go
│       ├── delim_string.go
│       ├── doc.go
│       ├── element_string.go
│       ├── error.go
│       ├── escape.go
│       ├── escape_test.go
│       ├── example_test.go
│       ├── examplefiles_test.go
│       ├── html.go
│       ├── html_test.go
│       ├── js.go
│       ├── js_test.go
│       ├── jsctx_string.go
│       ├── state_string.go
│       ├── template.go
│       ├── template_test.go
│       ├── transition.go
│       ├── transition_test.go
│       ├── url.go
│       ├── url_test.go
│       └── urlpart_string.go
├── image
│   ├── color
│   │   ├── color.go
│   │   ├── color_test.go
│   │   ├── palette
│   │   │   ├── gen.go
│   │   │   ├── generate.go
│   │   │   └── palette.go
│   │   ├── ycbcr.go
│   │   └── ycbcr_test.go
│   ├── decode_example_test.go
│   ├── decode_test.go
│   ├── draw
│   │   ├── bench_test.go
│   │   ├── clip_test.go
│   │   ├── draw.go
│   │   ├── draw_test.go
│   │   └── example_test.go
│   ├── format.go
│   ├── geom.go
│   ├── geom_test.go
│   ├── gif
│   │   ├── reader.go
│   │   ├── reader_test.go
│   │   ├── writer.go
│   │   └── writer_test.go
│   ├── image.go
│   ├── image_test.go
│   ├── internal
│   │   └── imageutil
│   │       ├── gen.go
│   │       ├── imageutil.go
│   │       └── impl.go
│   ├── jpeg
│   │   ├── dct_test.go
│   │   ├── fdct.go
│   │   ├── huffman.go
│   │   ├── idct.go
│   │   ├── reader.go
│   │   ├── reader_test.go
│   │   ├── scan.go
│   │   ├── writer.go
│   │   └── writer_test.go
│   ├── names.go
│   ├── png
│   │   ├── example_test.go
│   │   ├── fuzz.go
│   │   ├── paeth.go
│   │   ├── paeth_test.go
│   │   ├── reader.go
│   │   ├── reader_test.go
│   │   ├── testdata
│   │   │   ├── benchGray.png
│   │   │   ├── benchNRGBA-gradient.png
│   │   │   ├── benchNRGBA-opaque.png
│   │   │   ├── benchPaletted.png
│   │   │   ├── benchRGB-interlace.png
│   │   │   ├── benchRGB.png
│   │   │   ├── gray-gradient.interlaced.png
│   │   │   ├── gray-gradient.png
│   │   │   ├── invalid-crc32.png
│   │   │   ├── invalid-noend.png
│   │   │   ├── invalid-palette.png
│   │   │   ├── invalid-trunc.png
│   │   │   ├── invalid-zlib.png
│   │   │   └── pngsuite
│   │   │       ├── README
│   │   │       ├── README.original
│   │   │       ├── basn0g01-30.png
│   │   │       ├── basn0g01-30.sng
│   │   │       ├── basn0g01.png
│   │   │       ├── basn0g01.sng
│   │   │       ├── basn0g02-29.png
│   │   │       ├── basn0g02-29.sng
│   │   │       ├── basn0g02.png
│   │   │       ├── basn0g02.sng
│   │   │       ├── basn0g04-31.png
│   │   │       ├── basn0g04-31.sng
│   │   │       ├── basn0g04.png
│   │   │       ├── basn0g04.sng
│   │   │       ├── basn0g08.png
│   │   │       ├── basn0g08.sng
│   │   │       ├── basn0g16.png
│   │   │       ├── basn0g16.sng
│   │   │       ├── basn2c08.png
│   │   │       ├── basn2c08.sng
│   │   │       ├── basn2c16.png
│   │   │       ├── basn2c16.sng
│   │   │       ├── basn3p01.png
│   │   │       ├── basn3p01.sng
│   │   │       ├── basn3p02.png
│   │   │       ├── basn3p02.sng
│   │   │       ├── basn3p04-31i.png
│   │   │       ├── basn3p04-31i.sng
│   │   │       ├── basn3p04.png
│   │   │       ├── basn3p04.sng
│   │   │       ├── basn3p08-trns.png
│   │   │       ├── basn3p08-trns.sng
│   │   │       ├── basn3p08.png
│   │   │       ├── basn3p08.sng
│   │   │       ├── basn4a08.png
│   │   │       ├── basn4a08.sng
│   │   │       ├── basn4a16.png
│   │   │       ├── basn4a16.sng
│   │   │       ├── basn6a08.png
│   │   │       ├── basn6a08.sng
│   │   │       ├── basn6a16.png
│   │   │       ├── basn6a16.sng
│   │   │       ├── ftbbn0g01.png
│   │   │       ├── ftbbn0g01.sng
│   │   │       ├── ftbbn0g02.png
│   │   │       ├── ftbbn0g02.sng
│   │   │       ├── ftbbn0g04.png
│   │   │       ├── ftbbn0g04.sng
│   │   │       ├── ftbbn2c16.png
│   │   │       ├── ftbbn2c16.sng
│   │   │       ├── ftbbn3p08.png
│   │   │       ├── ftbbn3p08.sng
│   │   │       ├── ftbgn2c16.png
│   │   │       ├── ftbgn2c16.sng
│   │   │       ├── ftbgn3p08.png
│   │   │       ├── ftbgn3p08.sng
│   │   │       ├── ftbrn2c08.png
│   │   │       ├── ftbrn2c08.sng
│   │   │       ├── ftbwn0g16.png
│   │   │       ├── ftbwn0g16.sng
│   │   │       ├── ftbwn3p08.png
│   │   │       ├── ftbwn3p08.sng
│   │   │       ├── ftbyn3p08.png
│   │   │       ├── ftbyn3p08.sng
│   │   │       ├── ftp0n0g08.png
│   │   │       ├── ftp0n0g08.sng
│   │   │       ├── ftp0n2c08.png
│   │   │       ├── ftp0n2c08.sng
│   │   │       ├── ftp0n3p08.png
│   │   │       ├── ftp0n3p08.sng
│   │   │       ├── ftp1n3p08.png
│   │   │       └── ftp1n3p08.sng
│   │   ├── writer.go
│   │   └── writer_test.go
│   ├── testdata
│   │   ├── triangle-001.gif
│   │   ├── video-001.221212.jpeg
│   │   ├── video-001.221212.png
│   │   ├── video-001.5bpp.gif
│   │   ├── video-001.cmyk.jpeg
│   │   ├── video-001.cmyk.png
│   │   ├── video-001.gif
│   │   ├── video-001.interlaced.gif
│   │   ├── video-001.jpeg
│   │   ├── video-001.png
│   │   ├── video-001.progressive.jpeg
│   │   ├── video-001.progressive.truncated.jpeg
│   │   ├── video-001.progressive.truncated.png
│   │   ├── video-001.q50.410.jpeg
│   │   ├── video-001.q50.410.progressive.jpeg
│   │   ├── video-001.q50.411.jpeg
│   │   ├── video-001.q50.411.progressive.jpeg
│   │   ├── video-001.q50.420.jpeg
│   │   ├── video-001.q50.420.progressive.jpeg
│   │   ├── video-001.q50.422.jpeg
│   │   ├── video-001.q50.422.progressive.jpeg
│   │   ├── video-001.q50.440.jpeg
│   │   ├── video-001.q50.440.progressive.jpeg
│   │   ├── video-001.q50.444.jpeg
│   │   ├── video-001.q50.444.progressive.jpeg
│   │   ├── video-001.rgb.jpeg
│   │   ├── video-001.rgb.png
│   │   ├── video-001.separate.dc.progression.jpeg
│   │   ├── video-001.separate.dc.progression.progressive.jpeg
│   │   ├── video-005.gray.gif
│   │   ├── video-005.gray.jpeg
│   │   ├── video-005.gray.png
│   │   ├── video-005.gray.q50.2x2.jpeg
│   │   ├── video-005.gray.q50.2x2.progressive.jpeg
│   │   ├── video-005.gray.q50.jpeg
│   │   └── video-005.gray.q50.progressive.jpeg
│   ├── ycbcr.go
│   └── ycbcr_test.go
├── index
│   └── suffixarray
│       ├── example_test.go
│       ├── gen.go
│       ├── sais.go
│       ├── sais2.go
│       ├── suffixarray.go
│       └── suffixarray_test.go
├── internal
│   ├── bytealg
│   │   ├── bytealg.go
│   │   ├── compare_386.s
│   │   ├── compare_amd64.s
│   │   ├── compare_amd64p32.s
│   │   ├── compare_arm.s
│   │   ├── compare_arm64.s
│   │   ├── compare_generic.go
│   │   ├── compare_mipsx.s
│   │   ├── compare_native.go
│   │   ├── compare_ppc64x.s
│   │   ├── compare_s390x.s
│   │   ├── compare_wasm.s
│   │   ├── count_amd64.s
│   │   ├── count_arm.s
│   │   ├── count_arm64.s
│   │   ├── count_generic.go
│   │   ├── count_native.go
│   │   ├── count_ppc64x.s
│   │   ├── equal_386.s
│   │   ├── equal_amd64.s
│   │   ├── equal_amd64p32.s
│   │   ├── equal_arm.s
│   │   ├── equal_arm64.s
│   │   ├── equal_generic.go
│   │   ├── equal_mips64x.s
│   │   ├── equal_mipsx.s
│   │   ├── equal_native.go
│   │   ├── equal_ppc64x.s
│   │   ├── equal_s390x.s
│   │   ├── equal_wasm.s
│   │   ├── index_amd64.go
│   │   ├── index_amd64.s
│   │   ├── index_arm64.go
│   │   ├── index_arm64.s
│   │   ├── index_generic.go
│   │   ├── index_native.go
│   │   ├── index_s390x.go
│   │   ├── index_s390x.s
│   │   ├── indexbyte_386.s
│   │   ├── indexbyte_amd64.s
│   │   ├── indexbyte_amd64p32.s
│   │   ├── indexbyte_arm.s
│   │   ├── indexbyte_arm64.s
│   │   ├── indexbyte_generic.go
│   │   ├── indexbyte_mips64x.s
│   │   ├── indexbyte_mipsx.s
│   │   ├── indexbyte_native.go
│   │   ├── indexbyte_ppc64x.s
│   │   ├── indexbyte_s390x.s
│   │   └── indexbyte_wasm.s
│   ├── cfg
│   │   └── cfg.go
│   ├── cpu
│   │   ├── cpu.go
│   │   ├── cpu_386.go
│   │   ├── cpu_amd64.go
│   │   ├── cpu_amd64p32.go
│   │   ├── cpu_arm.go
│   │   ├── cpu_arm64.go
│   │   ├── cpu_mips.go
│   │   ├── cpu_mips64.go
│   │   ├── cpu_mips64le.go
│   │   ├── cpu_mipsle.go
│   │   ├── cpu_no_init.go
│   │   ├── cpu_ppc64x.go
│   │   ├── cpu_s390x.go
│   │   ├── cpu_s390x.s
│   │   ├── cpu_s390x_test.go
│   │   ├── cpu_test.go
│   │   ├── cpu_wasm.go
│   │   ├── cpu_x86.go
│   │   ├── cpu_x86.s
│   │   ├── cpu_x86_test.go
│   │   └── export_test.go
│   ├── fmtsort
│   │   ├── export_test.go
│   │   ├── sort.go
│   │   └── sort_test.go
│   ├── goroot
│   │   ├── gc.go
│   │   └── gccgo.go
│   ├── goversion
│   │   └── goversion.go
│   ├── lazyregexp
│   │   └── lazyre.go
│   ├── lazytemplate
│   │   └── lazytemplate.go
│   ├── nettrace
│   │   └── nettrace.go
│   ├── oserror
│   │   └── errors.go
│   ├── poll
│   │   ├── errno_unix.go
│   │   ├── errno_windows.go
│   │   ├── error_linux_test.go
│   │   ├── error_stub_test.go
│   │   ├── error_test.go
│   │   ├── export_posix_test.go
│   │   ├── export_test.go
│   │   ├── export_windows_test.go
│   │   ├── fd.go
│   │   ├── fd_fsync_darwin.go
│   │   ├── fd_fsync_posix.go
│   │   ├── fd_fsync_windows.go
│   │   ├── fd_io_plan9.go
│   │   ├── fd_mutex.go
│   │   ├── fd_mutex_test.go
│   │   ├── fd_opendir_darwin.go
│   │   ├── fd_plan9.go
│   │   ├── fd_poll_nacljs.go
│   │   ├── fd_poll_runtime.go
│   │   ├── fd_posix.go
│   │   ├── fd_posix_test.go
│   │   ├── fd_unix.go
│   │   ├── fd_windows.go
│   │   ├── fd_windows_test.go
│   │   ├── fd_writev_darwin.go
│   │   ├── fd_writev_unix.go
│   │   ├── hook_cloexec.go
│   │   ├── hook_unix.go
│   │   ├── hook_windows.go
│   │   ├── read_test.go
│   │   ├── sendfile_bsd.go
│   │   ├── sendfile_linux.go
│   │   ├── sendfile_solaris.go
│   │   ├── sendfile_windows.go
│   │   ├── sock_cloexec.go
│   │   ├── sockopt.go
│   │   ├── sockopt_linux.go
│   │   ├── sockopt_unix.go
│   │   ├── sockopt_windows.go
│   │   ├── sockoptip.go
│   │   ├── splice_linux.go
│   │   ├── strconv.go
│   │   ├── sys_cloexec.go
│   │   ├── writev.go
│   │   └── writev_test.go
│   ├── race
│   │   ├── doc.go
│   │   ├── norace.go
│   │   └── race.go
│   ├── reflectlite
│   │   ├── all_test.go
│   │   ├── asm.s
│   │   ├── export_test.go
│   │   ├── set_test.go
│   │   ├── swapper.go
│   │   ├── tostring_test.go
│   │   ├── type.go
│   │   └── value.go
│   ├── singleflight
│   │   ├── singleflight.go
│   │   └── singleflight_test.go
│   ├── syscall
│   │   ├── unix
│   │   │   ├── asm_aix_ppc64.s
│   │   │   ├── asm_solaris.s
│   │   │   ├── at.go
│   │   │   ├── at_aix.go
│   │   │   ├── at_darwin.go
│   │   │   ├── at_freebsd.go
│   │   │   ├── at_libc.go
│   │   │   ├── at_solaris.go
│   │   │   ├── at_sysnum_darwin.go
│   │   │   ├── at_sysnum_dragonfly.go
│   │   │   ├── at_sysnum_fstatat64_linux.go
│   │   │   ├── at_sysnum_fstatat_linux.go
│   │   │   ├── at_sysnum_linux.go
│   │   │   ├── at_sysnum_netbsd.go
│   │   │   ├── at_sysnum_newfstatat_linux.go
│   │   │   ├── at_sysnum_openbsd.go
│   │   │   ├── getentropy_openbsd.go
│   │   │   ├── getrandom_freebsd.go
│   │   │   ├── getrandom_linux.go
│   │   │   ├── getrandom_linux_386.go
│   │   │   ├── getrandom_linux_amd64.go
│   │   │   ├── getrandom_linux_arm.go
│   │   │   ├── getrandom_linux_generic.go
│   │   │   ├── getrandom_linux_mips64x.go
│   │   │   ├── getrandom_linux_mipsx.go
│   │   │   ├── getrandom_linux_ppc64x.go
│   │   │   ├── getrandom_linux_s390x.go
│   │   │   ├── ioctl_aix.go
│   │   │   ├── nonblocking.go
│   │   │   ├── nonblocking_darwin.go
│   │   │   ├── nonblocking_js.go
│   │   │   └── nonblocking_nacl.go
│   │   └── windows
│   │       ├── exec_windows_test.go
│   │       ├── mksyscall.go
│   │       ├── psapi_windows.go
│   │       ├── registry
│   │       │   ├── export_test.go
│   │       │   ├── key.go
│   │       │   ├── mksyscall.go
│   │       │   ├── registry_test.go
│   │       │   ├── syscall.go
│   │       │   ├── value.go
│   │       │   └── zsyscall_windows.go
│   │       ├── reparse_windows.go
│   │       ├── security_windows.go
│   │       ├── symlink_windows.go
│   │       ├── syscall_windows.go
│   │       ├── sysdll
│   │       │   └── sysdll.go
│   │       └── zsyscall_windows.go
│   ├── testenv
│   │   ├── testenv.go
│   │   ├── testenv_cgo.go
│   │   ├── testenv_notwin.go
│   │   └── testenv_windows.go
│   ├── testlog
│   │   └── log.go
│   ├── trace
│   │   ├── gc.go
│   │   ├── gc_test.go
│   │   ├── goroutines.go
│   │   ├── mkcanned.bash
│   │   ├── mud.go
│   │   ├── mud_test.go
│   │   ├── order.go
│   │   ├── parser.go
│   │   ├── parser_test.go
│   │   ├── testdata
│   │   │   ├── http_1_10_good
│   │   │   ├── http_1_11_good
│   │   │   ├── http_1_5_good
│   │   │   ├── http_1_7_good
│   │   │   ├── http_1_9_good
│   │   │   ├── stress_1_10_good
│   │   │   ├── stress_1_11_good
│   │   │   ├── stress_1_5_good
│   │   │   ├── stress_1_5_unordered
│   │   │   ├── stress_1_7_good
│   │   │   ├── stress_1_9_good
│   │   │   ├── stress_start_stop_1_10_good
│   │   │   ├── stress_start_stop_1_11_good
│   │   │   ├── stress_start_stop_1_5_good
│   │   │   ├── stress_start_stop_1_7_good
│   │   │   ├── stress_start_stop_1_9_good
│   │   │   └── user_task_span_1_11_good
│   │   └── writer.go
│   └── xcoff
│       ├── ar.go
│       ├── ar_test.go
│       ├── file.go
│       ├── file_test.go
│       ├── testdata
│       │   ├── bigar-empty
│       │   ├── bigar-ppc64
│       │   ├── gcc-ppc32-aix-dwarf2-exec
│       │   ├── gcc-ppc64-aix-dwarf2-exec
│       │   ├── hello.c
│       │   ├── printbye.c
│       │   └── printhello.c
│       └── xcoff.go
├── io
│   ├── example_test.go
│   ├── io.go
│   ├── io_test.go
│   ├── ioutil
│   │   ├── example_test.go
│   │   ├── ioutil.go
│   │   ├── ioutil_test.go
│   │   ├── tempfile.go
│   │   ├── tempfile_test.go
│   │   └── testdata
│   │       └── hello
│   ├── multi.go
│   ├── multi_test.go
│   ├── pipe.go
│   └── pipe_test.go
├── iostest.bash
├── log
│   ├── example_test.go
│   ├── log.go
│   ├── log_test.go
│   └── syslog
│       ├── doc.go
│       ├── example_test.go
│       ├── syslog.go
│       ├── syslog_test.go
│       └── syslog_unix.go
├── make.bash
├── make.bat
├── make.rc
├── math
│   ├── abs.go
│   ├── acos_s390x.s
│   ├── acosh.go
│   ├── acosh_s390x.s
│   ├── all_test.go
│   ├── arith_s390x.go
│   ├── arith_s390x_test.go
│   ├── asin.go
│   ├── asin_386.s
│   ├── asin_s390x.s
│   ├── asinh.go
│   ├── asinh_s390x.s
│   ├── atan.go
│   ├── atan2.go
│   ├── atan2_386.s
│   ├── atan2_s390x.s
│   ├── atan_386.s
│   ├── atan_s390x.s
│   ├── atanh.go
│   ├── atanh_s390x.s
│   ├── big
│   │   ├── accuracy_string.go
│   │   ├── arith.go
│   │   ├── arith_386.s
│   │   ├── arith_amd64.go
│   │   ├── arith_amd64.s
│   │   ├── arith_amd64p32.s
│   │   ├── arith_arm.s
│   │   ├── arith_arm64.s
│   │   ├── arith_decl.go
│   │   ├── arith_decl_pure.go
│   │   ├── arith_decl_s390x.go
│   │   ├── arith_mips64x.s
│   │   ├── arith_mipsx.s
│   │   ├── arith_ppc64x.s
│   │   ├── arith_s390x.s
│   │   ├── arith_s390x_test.go
│   │   ├── arith_test.go
│   │   ├── arith_wasm.s
│   │   ├── bits_test.go
│   │   ├── calibrate_test.go
│   │   ├── decimal.go
│   │   ├── decimal_test.go
│   │   ├── doc.go
│   │   ├── example_rat_test.go
│   │   ├── example_test.go
│   │   ├── float.go
│   │   ├── float_test.go
│   │   ├── floatconv.go
│   │   ├── floatconv_test.go
│   │   ├── floatexample_test.go
│   │   ├── floatmarsh.go
│   │   ├── floatmarsh_test.go
│   │   ├── ftoa.go
│   │   ├── gcd_test.go
│   │   ├── hilbert_test.go
│   │   ├── int.go
│   │   ├── int_test.go
│   │   ├── intconv.go
│   │   ├── intconv_test.go
│   │   ├── intmarsh.go
│   │   ├── intmarsh_test.go
│   │   ├── nat.go
│   │   ├── nat_test.go
│   │   ├── natconv.go
│   │   ├── natconv_test.go
│   │   ├── prime.go
│   │   ├── prime_test.go
│   │   ├── rat.go
│   │   ├── rat_test.go
│   │   ├── ratconv.go
│   │   ├── ratconv_test.go
│   │   ├── ratmarsh.go
│   │   ├── ratmarsh_test.go
│   │   ├── roundingmode_string.go
│   │   ├── sqrt.go
│   │   └── sqrt_test.go
│   ├── bits
│   │   ├── bits.go
│   │   ├── bits_errors.go
│   │   ├── bits_errors_bootstrap.go
│   │   ├── bits_tables.go
│   │   ├── bits_test.go
│   │   ├── example_test.go
│   │   ├── export_test.go
│   │   ├── make_examples.go
│   │   └── make_tables.go
│   ├── bits.go
│   ├── cbrt.go
│   ├── cbrt_s390x.s
│   ├── cmplx
│   │   ├── abs.go
│   │   ├── asin.go
│   │   ├── cmath_test.go
│   │   ├── conj.go
│   │   ├── example_test.go
│   │   ├── exp.go
│   │   ├── isinf.go
│   │   ├── isnan.go
│   │   ├── log.go
│   │   ├── phase.go
│   │   ├── polar.go
│   │   ├── pow.go
│   │   ├── rect.go
│   │   ├── sin.go
│   │   ├── sqrt.go
│   │   └── tan.go
│   ├── const.go
│   ├── copysign.go
│   ├── cosh_s390x.s
│   ├── dim.go
│   ├── dim_amd64.s
│   ├── dim_amd64p32.s
│   ├── dim_arm64.s
│   ├── dim_s390x.s
│   ├── erf.go
│   ├── erf_s390x.s
│   ├── erfc_s390x.s
│   ├── erfinv.go
│   ├── example_test.go
│   ├── exp.go
│   ├── exp2_386.s
│   ├── exp_amd64.s
│   ├── exp_amd64p32.s
│   ├── exp_arm64.s
│   ├── exp_asm.go
│   ├── exp_s390x.s
│   ├── expm1.go
│   ├── expm1_386.s
│   ├── expm1_s390x.s
│   ├── export_s390x_test.go
│   ├── export_test.go
│   ├── floor.go
│   ├── floor_386.s
│   ├── floor_amd64.s
│   ├── floor_amd64p32.s
│   ├── floor_arm64.s
│   ├── floor_ppc64x.s
│   ├── floor_s390x.s
│   ├── floor_wasm.s
│   ├── frexp.go
│   ├── frexp_386.s
│   ├── gamma.go
│   ├── huge_test.go
│   ├── hypot.go
│   ├── hypot_386.s
│   ├── hypot_amd64.s
│   ├── hypot_amd64p32.s
│   ├── j0.go
│   ├── j1.go
│   ├── jn.go
│   ├── ldexp.go
│   ├── ldexp_386.s
│   ├── lgamma.go
│   ├── log.go
│   ├── log10.go
│   ├── log10_386.s
│   ├── log10_s390x.s
│   ├── log1p.go
│   ├── log1p_386.s
│   ├── log1p_s390x.s
│   ├── log_386.s
│   ├── log_amd64.s
│   ├── log_amd64p32.s
│   ├── log_s390x.s
│   ├── logb.go
│   ├── mod.go
│   ├── mod_386.s
│   ├── modf.go
│   ├── modf_386.s
│   ├── modf_arm64.s
│   ├── modf_ppc64x.s
│   ├── nextafter.go
│   ├── pow.go
│   ├── pow10.go
│   ├── pow_s390x.s
│   ├── rand
│   │   ├── example_test.go
│   │   ├── exp.go
│   │   ├── gen_cooked.go
│   │   ├── normal.go
│   │   ├── race_test.go
│   │   ├── rand.go
│   │   ├── rand_test.go
│   │   ├── regress_test.go
│   │   ├── rng.go
│   │   └── zipf.go
│   ├── remainder.go
│   ├── remainder_386.s
│   ├── signbit.go
│   ├── sin.go
│   ├── sin_s390x.s
│   ├── sincos.go
│   ├── sinh.go
│   ├── sinh_s390x.s
│   ├── sqrt.go
│   ├── sqrt_386.s
│   ├── sqrt_amd64.s
│   ├── sqrt_amd64p32.s
│   ├── sqrt_arm.s
│   ├── sqrt_arm64.s
│   ├── sqrt_mipsx.s
│   ├── sqrt_ppc64x.s
│   ├── sqrt_s390x.s
│   ├── sqrt_wasm.s
│   ├── stubs_386.s
│   ├── stubs_amd64.s
│   ├── stubs_amd64p32.s
│   ├── stubs_arm.s
│   ├── stubs_arm64.s
│   ├── stubs_mips64x.s
│   ├── stubs_mipsx.s
│   ├── stubs_ppc64x.s
│   ├── stubs_s390x.s
│   ├── stubs_wasm.s
│   ├── tan.go
│   ├── tan_s390x.s
│   ├── tanh.go
│   ├── tanh_s390x.s
│   ├── trig_reduce.go
│   └── unsafe.go
├── mime
│   ├── encodedword.go
│   ├── encodedword_test.go
│   ├── example_test.go
│   ├── grammar.go
│   ├── mediatype.go
│   ├── mediatype_test.go
│   ├── multipart
│   │   ├── example_test.go
│   │   ├── formdata.go
│   │   ├── formdata_test.go
│   │   ├── multipart.go
│   │   ├── multipart_test.go
│   │   ├── testdata
│   │   │   └── nested-mime
│   │   ├── writer.go
│   │   └── writer_test.go
│   ├── quotedprintable
│   │   ├── example_test.go
│   │   ├── reader.go
│   │   ├── reader_test.go
│   │   ├── writer.go
│   │   └── writer_test.go
│   ├── testdata
│   │   ├── test.types
│   │   └── test.types.plan9
│   ├── type.go
│   ├── type_dragonfly.go
│   ├── type_freebsd.go
│   ├── type_openbsd.go
│   ├── type_plan9.go
│   ├── type_test.go
│   ├── type_unix.go
│   └── type_windows.go
├── naclmake.bash
├── nacltest.bash
├── net
│   ├── addrselect.go
│   ├── addrselect_test.go
│   ├── cgo_aix.go
│   ├── cgo_android.go
│   ├── cgo_bsd.go
│   ├── cgo_linux.go
│   ├── cgo_netbsd.go
│   ├── cgo_openbsd.go
│   ├── cgo_resnew.go
│   ├── cgo_resold.go
│   ├── cgo_socknew.go
│   ├── cgo_sockold.go
│   ├── cgo_solaris.go
│   ├── cgo_stub.go
│   ├── cgo_unix.go
│   ├── cgo_unix_test.go
│   ├── cgo_windows.go
│   ├── conf.go
│   ├── conf_netcgo.go
│   ├── conf_test.go
│   ├── conn_test.go
│   ├── dial.go
│   ├── dial_test.go
│   ├── dial_unix_test.go
│   ├── dnsclient.go
│   ├── dnsclient_test.go
│   ├── dnsclient_unix.go
│   ├── dnsclient_unix_test.go
│   ├── dnsconfig_unix.go
│   ├── dnsconfig_unix_test.go
│   ├── dnsname_test.go
│   ├── error_nacl.go
│   ├── error_plan9.go
│   ├── error_plan9_test.go
│   ├── error_posix.go
│   ├── error_posix_test.go
│   ├── error_test.go
│   ├── error_unix.go
│   ├── error_unix_test.go
│   ├── error_windows.go
│   ├── error_windows_test.go
│   ├── example_test.go
│   ├── external_test.go
│   ├── fd_plan9.go
│   ├── fd_unix.go
│   ├── fd_windows.go
│   ├── file.go
│   ├── file_plan9.go
│   ├── file_stub.go
│   ├── file_test.go
│   ├── file_unix.go
│   ├── file_windows.go
│   ├── hook.go
│   ├── hook_plan9.go
│   ├── hook_unix.go
│   ├── hook_windows.go
│   ├── hosts.go
│   ├── hosts_test.go
│   ├── http
│   │   ├── cgi
│   │   │   ├── child.go
│   │   │   ├── child_test.go
│   │   │   ├── host.go
│   │   │   ├── host_test.go
│   │   │   ├── matryoshka_test.go
│   │   │   ├── plan9_test.go
│   │   │   ├── posix_test.go
│   │   │   └── testdata
│   │   │       └── test.cgi
│   │   ├── client.go
│   │   ├── client_test.go
│   │   ├── clientserver_test.go
│   │   ├── clone.go
│   │   ├── cookie.go
│   │   ├── cookie_test.go
│   │   ├── cookiejar
│   │   │   ├── dummy_publicsuffix_test.go
│   │   │   ├── example_test.go
│   │   │   ├── jar.go
│   │   │   ├── jar_test.go
│   │   │   ├── punycode.go
│   │   │   └── punycode_test.go
│   │   ├── doc.go
│   │   ├── example_filesystem_test.go
│   │   ├── example_handle_test.go
│   │   ├── example_test.go
│   │   ├── export_test.go
│   │   ├── fcgi
│   │   │   ├── child.go
│   │   │   ├── fcgi.go
│   │   │   └── fcgi_test.go
│   │   ├── filetransport.go
│   │   ├── filetransport_test.go
│   │   ├── fs.go
│   │   ├── fs_test.go
│   │   ├── h2_bundle.go
│   │   ├── header.go
│   │   ├── header_test.go
│   │   ├── http.go
│   │   ├── http_test.go
│   │   ├── httptest
│   │   │   ├── example_test.go
│   │   │   ├── httptest.go
│   │   │   ├── httptest_test.go
│   │   │   ├── recorder.go
│   │   │   ├── recorder_test.go
│   │   │   ├── server.go
│   │   │   └── server_test.go
│   │   ├── httptrace
│   │   │   ├── example_test.go
│   │   │   ├── trace.go
│   │   │   └── trace_test.go
│   │   ├── httputil
│   │   │   ├── dump.go
│   │   │   ├── dump_test.go
│   │   │   ├── example_test.go
│   │   │   ├── httputil.go
│   │   │   ├── persist.go
│   │   │   ├── reverseproxy.go
│   │   │   └── reverseproxy_test.go
│   │   ├── internal
│   │   │   ├── chunked.go
│   │   │   ├── chunked_test.go
│   │   │   └── testcert.go
│   │   ├── jar.go
│   │   ├── main_test.go
│   │   ├── method.go
│   │   ├── npn_test.go
│   │   ├── pprof
│   │   │   ├── pprof.go
│   │   │   └── pprof_test.go
│   │   ├── proxy_test.go
│   │   ├── range_test.go
│   │   ├── readrequest_test.go
│   │   ├── request.go
│   │   ├── request_test.go
│   │   ├── requestwrite_test.go
│   │   ├── response.go
│   │   ├── response_test.go
│   │   ├── responsewrite_test.go
│   │   ├── roundtrip.go
│   │   ├── roundtrip_js.go
│   │   ├── serve_test.go
│   │   ├── server.go
│   │   ├── server_test.go
│   │   ├── sniff.go
│   │   ├── sniff_test.go
│   │   ├── socks_bundle.go
│   │   ├── status.go
│   │   ├── testdata
│   │   │   ├── file
│   │   │   ├── index.html
│   │   │   └── style.css
│   │   ├── transfer.go
│   │   ├── transfer_test.go
│   │   ├── transport.go
│   │   ├── transport_internal_test.go
│   │   ├── transport_test.go
│   │   └── triv.go
│   ├── interface.go
│   ├── interface_aix.go
│   ├── interface_bsd.go
│   ├── interface_bsd_test.go
│   ├── interface_bsdvar.go
│   ├── interface_darwin.go
│   ├── interface_freebsd.go
│   ├── interface_linux.go
│   ├── interface_linux_test.go
│   ├── interface_plan9.go
│   ├── interface_solaris.go
│   ├── interface_stub.go
│   ├── interface_test.go
│   ├── interface_unix_test.go
│   ├── interface_windows.go
│   ├── internal
│   │   └── socktest
│   │       ├── main_test.go
│   │       ├── main_unix_test.go
│   │       ├── main_windows_test.go
│   │       ├── switch.go
│   │       ├── switch_posix.go
│   │       ├── switch_stub.go
│   │       ├── switch_unix.go
│   │       ├── switch_windows.go
│   │       ├── sys_cloexec.go
│   │       ├── sys_unix.go
│   │       └── sys_windows.go
│   ├── ip.go
│   ├── ip_test.go
│   ├── iprawsock.go
│   ├── iprawsock_plan9.go
│   ├── iprawsock_posix.go
│   ├── iprawsock_test.go
│   ├── ipsock.go
│   ├── ipsock_plan9.go
│   ├── ipsock_posix.go
│   ├── ipsock_test.go
│   ├── listen_test.go
│   ├── lookup.go
│   ├── lookup_fake.go
│   ├── lookup_plan9.go
│   ├── lookup_test.go
│   ├── lookup_unix.go
│   ├── lookup_windows.go
│   ├── lookup_windows_test.go
│   ├── mac.go
│   ├── mac_test.go
│   ├── mail
│   │   ├── example_test.go
│   │   ├── message.go
│   │   └── message_test.go
│   ├── main_cloexec_test.go
│   ├── main_conf_test.go
│   ├── main_noconf_test.go
│   ├── main_plan9_test.go
│   ├── main_posix_test.go
│   ├── main_test.go
│   ├── main_unix_test.go
│   ├── main_windows_test.go
│   ├── mockserver_test.go
│   ├── net.go
│   ├── net_fake.go
│   ├── net_test.go
│   ├── net_windows_test.go
│   ├── netgo_unix_test.go
│   ├── nss.go
│   ├── nss_test.go
│   ├── packetconn_test.go
│   ├── parse.go
│   ├── parse_test.go
│   ├── pipe.go
│   ├── pipe_test.go
│   ├── platform_test.go
│   ├── port.go
│   ├── port_test.go
│   ├── port_unix.go
│   ├── protoconn_test.go
│   ├── rawconn.go
│   ├── rawconn_stub_test.go
│   ├── rawconn_test.go
│   ├── rawconn_unix_test.go
│   ├── rawconn_windows_test.go
│   ├── rpc
│   │   ├── client.go
│   │   ├── client_test.go
│   │   ├── debug.go
│   │   ├── jsonrpc
│   │   │   ├── all_test.go
│   │   │   ├── client.go
│   │   │   └── server.go
│   │   ├── server.go
│   │   └── server_test.go
│   ├── sendfile_linux.go
│   ├── sendfile_stub.go
│   ├── sendfile_test.go
│   ├── sendfile_unix_alt.go
│   ├── sendfile_windows.go
│   ├── server_test.go
│   ├── smtp
│   │   ├── auth.go
│   │   ├── example_test.go
│   │   ├── smtp.go
│   │   └── smtp_test.go
│   ├── sock_bsd.go
│   ├── sock_cloexec.go
│   ├── sock_linux.go
│   ├── sock_plan9.go
│   ├── sock_posix.go
│   ├── sock_stub.go
│   ├── sock_windows.go
│   ├── sockaddr_posix.go
│   ├── sockopt_aix.go
│   ├── sockopt_bsd.go
│   ├── sockopt_linux.go
│   ├── sockopt_plan9.go
│   ├── sockopt_posix.go
│   ├── sockopt_solaris.go
│   ├── sockopt_stub.go
│   ├── sockopt_windows.go
│   ├── sockoptip_bsdvar.go
│   ├── sockoptip_linux.go
│   ├── sockoptip_posix.go
│   ├── sockoptip_stub.go
│   ├── sockoptip_windows.go
│   ├── splice_linux.go
│   ├── splice_stub.go
│   ├── splice_test.go
│   ├── sys_cloexec.go
│   ├── tcpsock.go
│   ├── tcpsock_plan9.go
│   ├── tcpsock_posix.go
│   ├── tcpsock_test.go
│   ├── tcpsock_unix_test.go
│   ├── tcpsockopt_darwin.go
│   ├── tcpsockopt_dragonfly.go
│   ├── tcpsockopt_openbsd.go
│   ├── tcpsockopt_plan9.go
│   ├── tcpsockopt_posix.go
│   ├── tcpsockopt_solaris.go
│   ├── tcpsockopt_stub.go
│   ├── tcpsockopt_unix.go
│   ├── tcpsockopt_windows.go
│   ├── testdata
│   │   ├── case-hosts
│   │   ├── domain-resolv.conf
│   │   ├── empty-resolv.conf
│   │   ├── freebsd-usevc-resolv.conf
│   │   ├── hosts
│   │   ├── igmp
│   │   ├── igmp6
│   │   ├── invalid-ndots-resolv.conf
│   │   ├── ipv4-hosts
│   │   ├── ipv6-hosts
│   │   ├── large-ndots-resolv.conf
│   │   ├── linux-use-vc-resolv.conf
│   │   ├── negative-ndots-resolv.conf
│   │   ├── openbsd-resolv.conf
│   │   ├── openbsd-tcp-resolv.conf
│   │   ├── resolv.conf
│   │   ├── search-resolv.conf
│   │   ├── single-request-reopen-resolv.conf
│   │   ├── single-request-resolv.conf
│   │   └── singleline-hosts
│   ├── textproto
│   │   ├── header.go
│   │   ├── pipeline.go
│   │   ├── reader.go
│   │   ├── reader_test.go
│   │   ├── textproto.go
│   │   ├── writer.go
│   │   └── writer_test.go
│   ├── timeout_test.go
│   ├── udpsock.go
│   ├── udpsock_plan9.go
│   ├── udpsock_plan9_test.go
│   ├── udpsock_posix.go
│   ├── udpsock_test.go
│   ├── unixsock.go
│   ├── unixsock_linux_test.go
│   ├── unixsock_plan9.go
│   ├── unixsock_posix.go
│   ├── unixsock_test.go
│   ├── unixsock_windows_test.go
│   ├── url
│   │   ├── example_test.go
│   │   ├── url.go
│   │   └── url_test.go
│   ├── write_unix_test.go
│   ├── writev_test.go
│   └── writev_unix.go
├── os
│   ├── dir.go
│   ├── dir_darwin.go
│   ├── dir_plan9.go
│   ├── dir_unix.go
│   ├── dir_windows.go
│   ├── env.go
│   ├── env_default.go
│   ├── env_test.go
│   ├── env_unix_test.go
│   ├── env_windows.go
│   ├── error.go
│   ├── error_errno.go
│   ├── error_plan9.go
│   ├── error_posix.go
│   ├── error_test.go
│   ├── error_unix_test.go
│   ├── error_windows_test.go
│   ├── example_test.go
│   ├── exec
│   │   ├── bench_test.go
│   │   ├── env_test.go
│   │   ├── example_test.go
│   │   ├── exec.go
│   │   ├── exec_posix_test.go
│   │   ├── exec_test.go
│   │   ├── exec_unix.go
│   │   ├── exec_windows.go
│   │   ├── internal_test.go
│   │   ├── lp_js.go
│   │   ├── lp_plan9.go
│   │   ├── lp_test.go
│   │   ├── lp_unix.go
│   │   ├── lp_unix_test.go
│   │   ├── lp_windows.go
│   │   └── lp_windows_test.go
│   ├── exec.go
│   ├── exec_plan9.go
│   ├── exec_posix.go
│   ├── exec_unix.go
│   ├── exec_windows.go
│   ├── executable.go
│   ├── executable_darwin.go
│   ├── executable_freebsd.go
│   ├── executable_path.go
│   ├── executable_plan9.go
│   ├── executable_procfs.go
│   ├── executable_solaris.go
│   ├── executable_test.go
│   ├── executable_windows.go
│   ├── export_test.go
│   ├── export_unix_test.go
│   ├── export_windows_test.go
│   ├── fifo_test.go
│   ├── file.go
│   ├── file_plan9.go
│   ├── file_posix.go
│   ├── file_unix.go
│   ├── file_windows.go
│   ├── getwd.go
│   ├── getwd_darwin.go
│   ├── os_test.go
│   ├── os_unix_test.go
│   ├── os_windows_test.go
│   ├── path.go
│   ├── path_plan9.go
│   ├── path_test.go
│   ├── path_unix.go
│   ├── path_windows.go
│   ├── path_windows_test.go
│   ├── pipe2_bsd.go
│   ├── pipe_bsd.go
│   ├── pipe_linux.go
│   ├── pipe_test.go
│   ├── proc.go
│   ├── rawconn.go
│   ├── rawconn_test.go
│   ├── removeall_at.go
│   ├── removeall_noat.go
│   ├── removeall_test.go
│   ├── signal
│   │   ├── doc.go
│   │   ├── example_test.go
│   │   ├── internal
│   │   │   └── pty
│   │   │       └── pty.go
│   │   ├── sig.s
│   │   ├── signal.go
│   │   ├── signal_cgo_test.go
│   │   ├── signal_plan9.go
│   │   ├── signal_plan9_test.go
│   │   ├── signal_test.go
│   │   ├── signal_unix.go
│   │   └── signal_windows_test.go
│   ├── stat.go
│   ├── stat_aix.go
│   ├── stat_darwin.go
│   ├── stat_dragonfly.go
│   ├── stat_freebsd.go
│   ├── stat_linux.go
│   ├── stat_nacljs.go
│   ├── stat_netbsd.go
│   ├── stat_openbsd.go
│   ├── stat_plan9.go
│   ├── stat_solaris.go
│   ├── stat_test.go
│   ├── stat_unix.go
│   ├── stat_windows.go
│   ├── sticky_bsd.go
│   ├── sticky_notbsd.go
│   ├── str.go
│   ├── sys.go
│   ├── sys_aix.go
│   ├── sys_bsd.go
│   ├── sys_js.go
│   ├── sys_linux.go
│   ├── sys_nacl.go
│   ├── sys_plan9.go
│   ├── sys_solaris.go
│   ├── sys_unix.go
│   ├── sys_windows.go
│   ├── timeout_test.go
│   ├── types.go
│   ├── types_plan9.go
│   ├── types_unix.go
│   ├── types_windows.go
│   ├── user
│   │   ├── cgo_lookup_unix.go
│   │   ├── cgo_unix_test.go
│   │   ├── getgrouplist_darwin.go
│   │   ├── getgrouplist_unix.go
│   │   ├── listgroups_aix.go
│   │   ├── listgroups_solaris.go
│   │   ├── listgroups_unix.go
│   │   ├── lookup.go
│   │   ├── lookup_android.go
│   │   ├── lookup_plan9.go
│   │   ├── lookup_stubs.go
│   │   ├── lookup_unix.go
│   │   ├── lookup_unix_test.go
│   │   ├── lookup_windows.go
│   │   ├── user.go
│   │   └── user_test.go
│   ├── wait_unimp.go
│   ├── wait_wait6.go
│   └── wait_waitid.go
├── path
│   ├── example_test.go
│   ├── filepath
│   │   ├── example_test.go
│   │   ├── example_unix_test.go
│   │   ├── example_unix_walk_test.go
│   │   ├── export_test.go
│   │   ├── export_windows_test.go
│   │   ├── match.go
│   │   ├── match_test.go
│   │   ├── path.go
│   │   ├── path_plan9.go
│   │   ├── path_test.go
│   │   ├── path_unix.go
│   │   ├── path_windows.go
│   │   ├── path_windows_test.go
│   │   ├── symlink.go
│   │   ├── symlink_unix.go
│   │   └── symlink_windows.go
│   ├── match.go
│   ├── match_test.go
│   ├── path.go
│   └── path_test.go
├── plugin
│   ├── plugin.go
│   ├── plugin_dlopen.go
│   ├── plugin_stubs.go
│   └── plugin_test.go
├── race.bash
├── race.bat
├── reflect
│   ├── all_test.go
│   ├── asm_386.s
│   ├── asm_amd64.s
│   ├── asm_amd64p32.s
│   ├── asm_arm.s
│   ├── asm_arm64.s
│   ├── asm_mips64x.s
│   ├── asm_mipsx.s
│   ├── asm_ppc64x.s
│   ├── asm_s390x.s
│   ├── asm_wasm.s
│   ├── deepequal.go
│   ├── example_test.go
│   ├── export_test.go
│   ├── makefunc.go
│   ├── set_test.go
│   ├── swapper.go
│   ├── tostring_test.go
│   ├── type.go
│   └── value.go
├── regexp
│   ├── all_test.go
│   ├── backtrack.go
│   ├── example_test.go
│   ├── exec.go
│   ├── exec2_test.go
│   ├── exec_test.go
│   ├── find_test.go
│   ├── onepass.go
│   ├── onepass_test.go
│   ├── regexp.go
│   ├── syntax
│   │   ├── compile.go
│   │   ├── doc.go
│   │   ├── make_perl_groups.pl
│   │   ├── op_string.go
│   │   ├── parse.go
│   │   ├── parse_test.go
│   │   ├── perl_groups.go
│   │   ├── prog.go
│   │   ├── prog_test.go
│   │   ├── regexp.go
│   │   ├── simplify.go
│   │   └── simplify_test.go
│   └── testdata
│       ├── README
│       ├── basic.dat
│       ├── nullsubexpr.dat
│       ├── re2-exhaustive.txt.bz2
│       ├── re2-search.txt
│       ├── repetition.dat
│       └── testregex.c
├── run.bash
├── run.bat
├── run.rc
├── runtime
│   ├── HACKING.md
│   ├── Makefile
│   ├── alg.go
│   ├── asm.s
│   ├── asm_386.s
│   ├── asm_amd64.s
│   ├── asm_amd64p32.s
│   ├── asm_arm.s
│   ├── asm_arm64.s
│   ├── asm_mips64x.s
│   ├── asm_mipsx.s
│   ├── asm_ppc64x.h
│   ├── asm_ppc64x.s
│   ├── asm_s390x.s
│   ├── asm_wasm.s
│   ├── atomic_arm64.s
│   ├── atomic_mips64x.s
│   ├── atomic_mipsx.s
│   ├── atomic_pointer.go
│   ├── atomic_ppc64x.s
│   ├── auxv_none.go
│   ├── callers_test.go
│   ├── cgo
│   │   ├── asm_386.s
│   │   ├── asm_amd64.s
│   │   ├── asm_arm.s
│   │   ├── asm_arm64.s
│   │   ├── asm_mips64x.s
│   │   ├── asm_mipsx.s
│   │   ├── asm_nacl_amd64p32.s
│   │   ├── asm_ppc64x.s
│   │   ├── asm_s390x.s
│   │   ├── asm_wasm.s
│   │   ├── callbacks.go
│   │   ├── callbacks_aix.go
│   │   ├── callbacks_traceback.go
│   │   ├── cgo.go
│   │   ├── dragonfly.go
│   │   ├── freebsd.go
│   │   ├── gcc_386.S
│   │   ├── gcc_aix_ppc64.S
│   │   ├── gcc_aix_ppc64.c
│   │   ├── gcc_amd64.S
│   │   ├── gcc_android.c
│   │   ├── gcc_arm.S
│   │   ├── gcc_arm64.S
│   │   ├── gcc_context.c
│   │   ├── gcc_darwin_386.c
│   │   ├── gcc_darwin_amd64.c
│   │   ├── gcc_darwin_arm.c
│   │   ├── gcc_darwin_arm64.c
│   │   ├── gcc_dragonfly_amd64.c
│   │   ├── gcc_fatalf.c
│   │   ├── gcc_freebsd_386.c
│   │   ├── gcc_freebsd_amd64.c
│   │   ├── gcc_freebsd_arm.c
│   │   ├── gcc_freebsd_sigaction.c
│   │   ├── gcc_libinit.c
│   │   ├── gcc_libinit_windows.c
│   │   ├── gcc_linux_386.c
│   │   ├── gcc_linux_amd64.c
│   │   ├── gcc_linux_arm.c
│   │   ├── gcc_linux_arm64.c
│   │   ├── gcc_linux_mips64x.c
│   │   ├── gcc_linux_mipsx.c
│   │   ├── gcc_linux_ppc64x.S
│   │   ├── gcc_linux_s390x.c
│   │   ├── gcc_mips64x.S
│   │   ├── gcc_mipsx.S
│   │   ├── gcc_mmap.c
│   │   ├── gcc_netbsd_386.c
│   │   ├── gcc_netbsd_amd64.c
│   │   ├── gcc_netbsd_arm.c
│   │   ├── gcc_netbsd_arm64.c
│   │   ├── gcc_openbsd_386.c
│   │   ├── gcc_openbsd_amd64.c
│   │   ├── gcc_openbsd_arm.c
│   │   ├── gcc_openbsd_arm64.c
│   │   ├── gcc_ppc64x.c
│   │   ├── gcc_s390x.S
│   │   ├── gcc_setenv.c
│   │   ├── gcc_sigaction.c
│   │   ├── gcc_signal2_darwin_armx.c
│   │   ├── gcc_signal_darwin_armx.c
│   │   ├── gcc_signal_darwin_lldb.c
│   │   ├── gcc_solaris_amd64.c
│   │   ├── gcc_traceback.c
│   │   ├── gcc_util.c
│   │   ├── gcc_windows_386.c
│   │   ├── gcc_windows_amd64.c
│   │   ├── iscgo.go
│   │   ├── libcgo.h
│   │   ├── libcgo_unix.h
│   │   ├── mmap.go
│   │   ├── netbsd.go
│   │   ├── openbsd.go
│   │   ├── setenv.go
│   │   ├── sigaction.go
│   │   ├── signal_darwin_arm.s
│   │   ├── signal_darwin_arm64.s
│   │   └── signal_darwin_armx.go
│   ├── cgo.go
│   ├── cgo_mmap.go
│   ├── cgo_ppc64x.go
│   ├── cgo_sigaction.go
│   ├── cgocall.go
│   ├── cgocallback.go
│   ├── cgocheck.go
│   ├── chan.go
│   ├── chan_test.go
│   ├── chanbarrier_test.go
│   ├── closure_test.go
│   ├── compiler.go
│   ├── complex.go
│   ├── complex_test.go
│   ├── cpuflags.go
│   ├── cpuflags_amd64.go
│   ├── cpuprof.go
│   ├── cputicks.go
│   ├── crash_cgo_test.go
│   ├── crash_nonunix_test.go
│   ├── crash_test.go
│   ├── crash_unix_test.go
│   ├── debug
│   │   ├── debug.s
│   │   ├── garbage.go
│   │   ├── garbage_test.go
│   │   ├── heapdump_test.go
│   │   ├── mod.go
│   │   ├── stack.go
│   │   ├── stack_test.go
│   │   └── stubs.go
│   ├── debug.go
│   ├── debug_test.go
│   ├── debugcall.go
│   ├── debuglog.go
│   ├── debuglog_off.go
│   ├── debuglog_on.go
│   ├── debuglog_test.go
│   ├── defs1_linux.go
│   ├── defs1_netbsd_386.go
│   ├── defs1_netbsd_amd64.go
│   ├── defs1_netbsd_arm.go
│   ├── defs1_netbsd_arm64.go
│   ├── defs1_solaris_amd64.go
│   ├── defs2_linux.go
│   ├── defs3_linux.go
│   ├── defs_aix.go
│   ├── defs_aix_ppc64.go
│   ├── defs_arm_linux.go
│   ├── defs_darwin.go
│   ├── defs_darwin_386.go
│   ├── defs_darwin_amd64.go
│   ├── defs_darwin_arm.go
│   ├── defs_darwin_arm64.go
│   ├── defs_dragonfly.go
│   ├── defs_dragonfly_amd64.go
│   ├── defs_freebsd.go
│   ├── defs_freebsd_386.go
│   ├── defs_freebsd_amd64.go
│   ├── defs_freebsd_arm.go
│   ├── defs_linux.go
│   ├── defs_linux_386.go
│   ├── defs_linux_amd64.go
│   ├── defs_linux_arm.go
│   ├── defs_linux_arm64.go
│   ├── defs_linux_mips64x.go
│   ├── defs_linux_mipsx.go
│   ├── defs_linux_ppc64.go
│   ├── defs_linux_ppc64le.go
│   ├── defs_linux_s390x.go
│   ├── defs_nacl_386.go
│   ├── defs_nacl_amd64p32.go
│   ├── defs_nacl_arm.go
│   ├── defs_netbsd.go
│   ├── defs_netbsd_386.go
│   ├── defs_netbsd_amd64.go
│   ├── defs_netbsd_arm.go
│   ├── defs_openbsd.go
│   ├── defs_openbsd_386.go
│   ├── defs_openbsd_amd64.go
│   ├── defs_openbsd_arm.go
│   ├── defs_openbsd_arm64.go
│   ├── defs_plan9_386.go
│   ├── defs_plan9_amd64.go
│   ├── defs_plan9_arm.go
│   ├── defs_solaris.go
│   ├── defs_solaris_amd64.go
│   ├── defs_windows.go
│   ├── defs_windows_386.go
│   ├── defs_windows_amd64.go
│   ├── defs_windows_arm.go
│   ├── duff_386.s
│   ├── duff_amd64.s
│   ├── duff_arm.s
│   ├── duff_arm64.s
│   ├── duff_mips64x.s
│   ├── duff_ppc64x.s
│   ├── duff_s390x.s
│   ├── env_plan9.go
│   ├── env_posix.go
│   ├── env_test.go
│   ├── error.go
│   ├── example_test.go
│   ├── export_arm_test.go
│   ├── export_debug_test.go
│   ├── export_debuglog_test.go
│   ├── export_futex_test.go
│   ├── export_linux_test.go
│   ├── export_mmap_test.go
│   ├── export_test.go
│   ├── export_unix_test.go
│   ├── export_windows_test.go
│   ├── extern.go
│   ├── fastlog2.go
│   ├── fastlog2_test.go
│   ├── fastlog2table.go
│   ├── float.go
│   ├── funcdata.h
│   ├── futex_test.go
│   ├── gc_test.go
│   ├── gcinfo_test.go
│   ├── go_tls.h
│   ├── hash32.go
│   ├── hash64.go
│   ├── hash_test.go
│   ├── heapdump.go
│   ├── iface.go
│   ├── iface_test.go
│   ├── internal
│   │   ├── atomic
│   │   │   ├── asm_386.s
│   │   │   ├── asm_amd64.s
│   │   │   ├── asm_amd64p32.s
│   │   │   ├── asm_arm.s
│   │   │   ├── asm_arm64.s
│   │   │   ├── asm_mips64x.s
│   │   │   ├── asm_mipsx.s
│   │   │   ├── asm_ppc64x.s
│   │   │   ├── asm_s390x.s
│   │   │   ├── atomic_386.go
│   │   │   ├── atomic_amd64x.go
│   │   │   ├── atomic_arm.go
│   │   │   ├── atomic_arm64.go
│   │   │   ├── atomic_arm64.s
│   │   │   ├── atomic_mips64x.go
│   │   │   ├── atomic_mips64x.s
│   │   │   ├── atomic_mipsx.go
│   │   │   ├── atomic_mipsx.s
│   │   │   ├── atomic_ppc64x.go
│   │   │   ├── atomic_ppc64x.s
│   │   │   ├── atomic_s390x.go
│   │   │   ├── atomic_test.go
│   │   │   ├── atomic_wasm.go
│   │   │   ├── bench_test.go
│   │   │   ├── stubs.go
│   │   │   ├── sys_linux_arm.s
│   │   │   └── sys_nonlinux_arm.s
│   │   ├── math
│   │   │   ├── math.go
│   │   │   └── math_test.go
│   │   └── sys
│   │       ├── arch.go
│   │       ├── arch_386.go
│   │       ├── arch_amd64.go
│   │       ├── arch_amd64p32.go
│   │       ├── arch_arm.go
│   │       ├── arch_arm64.go
│   │       ├── arch_mips.go
│   │       ├── arch_mips64.go
│   │       ├── arch_mips64le.go
│   │       ├── arch_mipsle.go
│   │       ├── arch_ppc64.go
│   │       ├── arch_ppc64le.go
│   │       ├── arch_s390x.go
│   │       ├── arch_wasm.go
│   │       ├── gengoos.go
│   │       ├── intrinsics.go
│   │       ├── intrinsics_386.s
│   │       ├── intrinsics_stubs.go
│   │       ├── intrinsics_test.go
│   │       ├── stubs.go
│   │       ├── sys.go
│   │       ├── zgoarch_386.go
│   │       ├── zgoarch_amd64.go
│   │       ├── zgoarch_amd64p32.go
│   │       ├── zgoarch_arm.go
│   │       ├── zgoarch_arm64.go
│   │       ├── zgoarch_arm64be.go
│   │       ├── zgoarch_armbe.go
│   │       ├── zgoarch_mips.go
│   │       ├── zgoarch_mips64.go
│   │       ├── zgoarch_mips64le.go
│   │       ├── zgoarch_mips64p32.go
│   │       ├── zgoarch_mips64p32le.go
│   │       ├── zgoarch_mipsle.go
│   │       ├── zgoarch_ppc.go
│   │       ├── zgoarch_ppc64.go
│   │       ├── zgoarch_ppc64le.go
│   │       ├── zgoarch_riscv.go
│   │       ├── zgoarch_riscv64.go
│   │       ├── zgoarch_s390.go
│   │       ├── zgoarch_s390x.go
│   │       ├── zgoarch_sparc.go
│   │       ├── zgoarch_sparc64.go
│   │       ├── zgoarch_wasm.go
│   │       ├── zgoos_aix.go
│   │       ├── zgoos_android.go
│   │       ├── zgoos_darwin.go
│   │       ├── zgoos_dragonfly.go
│   │       ├── zgoos_freebsd.go
│   │       ├── zgoos_hurd.go
│   │       ├── zgoos_illumos.go
│   │       ├── zgoos_js.go
│   │       ├── zgoos_linux.go
│   │       ├── zgoos_nacl.go
│   │       ├── zgoos_netbsd.go
│   │       ├── zgoos_openbsd.go
│   │       ├── zgoos_plan9.go
│   │       ├── zgoos_solaris.go
│   │       ├── zgoos_windows.go
│   │       ├── zgoos_zos.go
│   │       └── zversion.go
│   ├── lfstack.go
│   ├── lfstack_32bit.go
│   ├── lfstack_64bit.go
│   ├── lfstack_test.go
│   ├── lock_futex.go
│   ├── lock_js.go
│   ├── lock_sema.go
│   ├── malloc.go
│   ├── malloc_test.go
│   ├── map.go
│   ├── map_benchmark_test.go
│   ├── map_fast32.go
│   ├── map_fast64.go
│   ├── map_faststr.go
│   ├── map_test.go
│   ├── mbarrier.go
│   ├── mbitmap.go
│   ├── mcache.go
│   ├── mcentral.go
│   ├── mem_aix.go
│   ├── mem_bsd.go
│   ├── mem_darwin.go
│   ├── mem_js.go
│   ├── mem_linux.go
│   ├── mem_plan9.go
│   ├── mem_windows.go
│   ├── memclr_386.s
│   ├── memclr_amd64.s
│   ├── memclr_amd64p32.s
│   ├── memclr_arm.s
│   ├── memclr_arm64.s
│   ├── memclr_mips64x.s
│   ├── memclr_mipsx.s
│   ├── memclr_plan9_386.s
│   ├── memclr_plan9_amd64.s
│   ├── memclr_ppc64x.s
│   ├── memclr_s390x.s
│   ├── memclr_wasm.s
│   ├── memmove_386.s
│   ├── memmove_amd64.s
│   ├── memmove_amd64p32.s
│   ├── memmove_arm.s
│   ├── memmove_arm64.s
│   ├── memmove_linux_amd64_test.go
│   ├── memmove_mips64x.s
│   ├── memmove_mipsx.s
│   ├── memmove_plan9_386.s
│   ├── memmove_plan9_amd64.s
│   ├── memmove_ppc64x.s
│   ├── memmove_s390x.s
│   ├── memmove_test.go
│   ├── memmove_wasm.s
│   ├── mfinal.go
│   ├── mfinal_test.go
│   ├── mfixalloc.go
│   ├── mgc.go
│   ├── mgclarge.go
│   ├── mgcmark.go
│   ├── mgcscavenge.go
│   ├── mgcstack.go
│   ├── mgcsweep.go
│   ├── mgcsweepbuf.go
│   ├── mgcwork.go
│   ├── mheap.go
│   ├── mkduff.go
│   ├── mkfastlog2table.go
│   ├── mknacl.sh
│   ├── mksizeclasses.go
│   ├── mmap.go
│   ├── mprof.go
│   ├── msan
│   │   └── msan.go
│   ├── msan.go
│   ├── msan0.go
│   ├── msan_amd64.s
│   ├── msan_arm64.s
│   ├── msize.go
│   ├── mstats.go
│   ├── mwbbuf.go
│   ├── net_plan9.go
│   ├── netpoll.go
│   ├── netpoll_aix.go
│   ├── netpoll_epoll.go
│   ├── netpoll_fake.go
│   ├── netpoll_kqueue.go
│   ├── netpoll_solaris.go
│   ├── netpoll_stub.go
│   ├── netpoll_windows.go
│   ├── norace_linux_test.go
│   ├── norace_test.go
│   ├── numcpu_freebsd_test.go
│   ├── os2_aix.go
│   ├── os2_freebsd.go
│   ├── os2_nacl.go
│   ├── os2_openbsd.go
│   ├── os2_plan9.go
│   ├── os2_solaris.go
│   ├── os3_plan9.go
│   ├── os3_solaris.go
│   ├── os_aix.go
│   ├── os_android.go
│   ├── os_darwin.go
│   ├── os_darwin_arm.go
│   ├── os_darwin_arm64.go
│   ├── os_dragonfly.go
│   ├── os_freebsd.go
│   ├── os_freebsd2.go
│   ├── os_freebsd_amd64.go
│   ├── os_freebsd_arm.go
│   ├── os_freebsd_noauxv.go
│   ├── os_js.go
│   ├── os_linux.go
│   ├── os_linux_arm.go
│   ├── os_linux_arm64.go
│   ├── os_linux_be64.go
│   ├── os_linux_generic.go
│   ├── os_linux_mips64x.go
│   ├── os_linux_mipsx.go
│   ├── os_linux_noauxv.go
│   ├── os_linux_novdso.go
│   ├── os_linux_ppc64x.go
│   ├── os_linux_s390x.go
│   ├── os_nacl.go
│   ├── os_nacl_arm.go
│   ├── os_netbsd.go
│   ├── os_netbsd_386.go
│   ├── os_netbsd_amd64.go
│   ├── os_netbsd_arm.go
│   ├── os_netbsd_arm64.go
│   ├── os_nonopenbsd.go
│   ├── os_openbsd.go
│   ├── os_openbsd_arm.go
│   ├── os_openbsd_arm64.go
│   ├── os_plan9.go
│   ├── os_plan9_arm.go
│   ├── os_solaris.go
│   ├── os_windows.go
│   ├── os_windows_arm.go
│   ├── panic.go
│   ├── panic32.go
│   ├── plugin.go
│   ├── pprof
│   │   ├── elf.go
│   │   ├── internal
│   │   │   └── profile
│   │   │       ├── encode.go
│   │   │       ├── filter.go
│   │   │       ├── legacy_profile.go
│   │   │       ├── profile.go
│   │   │       ├── profile_test.go
│   │   │       ├── proto.go
│   │   │       ├── proto_test.go
│   │   │       └── prune.go
│   │   ├── label.go
│   │   ├── label_test.go
│   │   ├── map.go
│   │   ├── mprof_test.go
│   │   ├── pprof.go
│   │   ├── pprof_test.go
│   │   ├── proto.go
│   │   ├── proto_test.go
│   │   ├── protobuf.go
│   │   ├── protomem.go
│   │   ├── protomem_test.go
│   │   ├── runtime.go
│   │   ├── runtime_test.go
│   │   └── testdata
│   │       ├── README
│   │       ├── mappingtest
│   │       │   └── main.go
│   │       ├── test32
│   │       ├── test32be
│   │       ├── test64
│   │       └── test64be
│   ├── print.go
│   ├── proc.go
│   ├── proc_runtime_test.go
│   ├── proc_test.go
│   ├── profbuf.go
│   ├── profbuf_test.go
│   ├── proflabel.go
│   ├── race
│   │   ├── README
│   │   ├── doc.go
│   │   ├── output_test.go
│   │   ├── race.go
│   │   ├── race_darwin_amd64.syso
│   │   ├── race_freebsd_amd64.syso
│   │   ├── race_linux_amd64.syso
│   │   ├── race_linux_arm64.syso
│   │   ├── race_linux_ppc64le.syso
│   │   ├── race_linux_test.go
│   │   ├── race_netbsd_amd64.syso
│   │   ├── race_test.go
│   │   ├── race_unix_test.go
│   │   ├── race_windows_amd64.syso
│   │   ├── race_windows_test.go
│   │   ├── sched_test.go
│   │   └── testdata
│   │       ├── atomic_test.go
│   │       ├── cgo_test.go
│   │       ├── cgo_test_main.go
│   │       ├── chan_test.go
│   │       ├── comp_test.go
│   │       ├── finalizer_test.go
│   │       ├── io_test.go
│   │       ├── issue12225_test.go
│   │       ├── issue12664_test.go
│   │       ├── issue13264_test.go
│   │       ├── map_test.go
│   │       ├── mop_test.go
│   │       ├── mutex_test.go
│   │       ├── pool_test.go
│   │       ├── reflect_test.go
│   │       ├── regression_test.go
│   │       ├── rwmutex_test.go
│   │       ├── select_test.go
│   │       ├── slice_test.go
│   │       ├── sync_test.go
│   │       └── waitgroup_test.go
│   ├── race.go
│   ├── race0.go
│   ├── race_amd64.s
│   ├── race_arm64.s
│   ├── race_ppc64le.s
│   ├── rand_test.go
│   ├── rdebug.go
│   ├── relax_stub.go
│   ├── rt0_aix_ppc64.s
│   ├── rt0_android_386.s
│   ├── rt0_android_amd64.s
│   ├── rt0_android_arm.s
│   ├── rt0_android_arm64.s
│   ├── rt0_darwin_386.s
│   ├── rt0_darwin_amd64.s
│   ├── rt0_darwin_arm.s
│   ├── rt0_darwin_arm64.s
│   ├── rt0_dragonfly_amd64.s
│   ├── rt0_freebsd_386.s
│   ├── rt0_freebsd_amd64.s
│   ├── rt0_freebsd_arm.s
│   ├── rt0_illumos_amd64.s
│   ├── rt0_js_wasm.s
│   ├── rt0_linux_386.s
│   ├── rt0_linux_amd64.s
│   ├── rt0_linux_arm.s
│   ├── rt0_linux_arm64.s
│   ├── rt0_linux_mips64x.s
│   ├── rt0_linux_mipsx.s
│   ├── rt0_linux_ppc64.s
│   ├── rt0_linux_ppc64le.s
│   ├── rt0_linux_s390x.s
│   ├── rt0_nacl_386.s
│   ├── rt0_nacl_amd64p32.s
│   ├── rt0_nacl_arm.s
│   ├── rt0_netbsd_386.s
│   ├── rt0_netbsd_amd64.s
│   ├── rt0_netbsd_arm.s
│   ├── rt0_netbsd_arm64.s
│   ├── rt0_openbsd_386.s
│   ├── rt0_openbsd_amd64.s
│   ├── rt0_openbsd_arm.s
│   ├── rt0_openbsd_arm64.s
│   ├── rt0_plan9_386.s
│   ├── rt0_plan9_amd64.s
│   ├── rt0_plan9_arm.s
│   ├── rt0_solaris_amd64.s
│   ├── rt0_windows_386.s
│   ├── rt0_windows_amd64.s
│   ├── rt0_windows_arm.s
│   ├── runtime-gdb.py
│   ├── runtime-gdb_test.go
│   ├── runtime-lldb_test.go
│   ├── runtime.go
│   ├── runtime1.go
│   ├── runtime2.go
│   ├── runtime_linux_test.go
│   ├── runtime_mmap_test.go
│   ├── runtime_test.go
│   ├── runtime_unix_test.go
│   ├── rwmutex.go
│   ├── rwmutex_test.go
│   ├── select.go
│   ├── sema.go
│   ├── semasleep_test.go
│   ├── sigaction.go
│   ├── signal_386.go
│   ├── signal_aix_ppc64.go
│   ├── signal_amd64x.go
│   ├── signal_arm.go
│   ├── signal_arm64.go
│   ├── signal_darwin.go
│   ├── signal_darwin_386.go
│   ├── signal_darwin_amd64.go
│   ├── signal_darwin_arm.go
│   ├── signal_darwin_arm64.go
│   ├── signal_dragonfly.go
│   ├── signal_dragonfly_amd64.go
│   ├── signal_freebsd.go
│   ├── signal_freebsd_386.go
│   ├── signal_freebsd_amd64.go
│   ├── signal_freebsd_arm.go
│   ├── signal_linux_386.go
│   ├── signal_linux_amd64.go
│   ├── signal_linux_arm.go
│   ├── signal_linux_arm64.go
│   ├── signal_linux_mips64x.go
│   ├── signal_linux_mipsx.go
│   ├── signal_linux_ppc64x.go
│   ├── signal_linux_s390x.go
│   ├── signal_mips64x.go
│   ├── signal_mipsx.go
│   ├── signal_nacl.go
│   ├── signal_nacl_386.go
│   ├── signal_nacl_amd64p32.go
│   ├── signal_nacl_arm.go
│   ├── signal_netbsd.go
│   ├── signal_netbsd_386.go
│   ├── signal_netbsd_amd64.go
│   ├── signal_netbsd_arm.go
│   ├── signal_netbsd_arm64.go
│   ├── signal_openbsd.go
│   ├── signal_openbsd_386.go
│   ├── signal_openbsd_amd64.go
│   ├── signal_openbsd_arm.go
│   ├── signal_openbsd_arm64.go
│   ├── signal_plan9.go
│   ├── signal_ppc64x.go
│   ├── signal_sighandler.go
│   ├── signal_solaris.go
│   ├── signal_solaris_amd64.go
│   ├── signal_unix.go
│   ├── signal_windows.go
│   ├── sigqueue.go
│   ├── sigqueue_note.go
│   ├── sigqueue_plan9.go
│   ├── sigtab_aix.go
│   ├── sigtab_linux_generic.go
│   ├── sigtab_linux_mipsx.go
│   ├── sizeclasses.go
│   ├── sizeof_test.go
│   ├── slice.go
│   ├── slice_test.go
│   ├── softfloat64.go
│   ├── softfloat64_test.go
│   ├── stack.go
│   ├── stack_test.go
│   ├── string.go
│   ├── string_test.go
│   ├── stubs.go
│   ├── stubs2.go
│   ├── stubs3.go
│   ├── stubs32.go
│   ├── stubs_386.go
│   ├── stubs_amd64x.go
│   ├── stubs_arm.go
│   ├── stubs_arm64.go
│   ├── stubs_linux.go
│   ├── stubs_mips64x.go
│   ├── stubs_mipsx.go
│   ├── stubs_nonlinux.go
│   ├── stubs_ppc64x.go
│   ├── stubs_s390x.go
│   ├── symtab.go
│   ├── symtab_test.go
│   ├── sys_aix_ppc64.s
│   ├── sys_arm.go
│   ├── sys_arm64.go
│   ├── sys_darwin.go
│   ├── sys_darwin_32.go
│   ├── sys_darwin_386.s
│   ├── sys_darwin_64.go
│   ├── sys_darwin_amd64.s
│   ├── sys_darwin_arm.s
│   ├── sys_darwin_arm64.s
│   ├── sys_dragonfly_amd64.s
│   ├── sys_freebsd_386.s
│   ├── sys_freebsd_amd64.s
│   ├── sys_freebsd_arm.s
│   ├── sys_linux_386.s
│   ├── sys_linux_amd64.s
│   ├── sys_linux_arm.s
│   ├── sys_linux_arm64.s
│   ├── sys_linux_mips64x.s
│   ├── sys_linux_mipsx.s
│   ├── sys_linux_ppc64x.s
│   ├── sys_linux_s390x.s
│   ├── sys_mips64x.go
│   ├── sys_mipsx.go
│   ├── sys_nacl_386.s
│   ├── sys_nacl_amd64p32.s
│   ├── sys_nacl_arm.s
│   ├── sys_netbsd_386.s
│   ├── sys_netbsd_amd64.s
│   ├── sys_netbsd_arm.s
│   ├── sys_netbsd_arm64.s
│   ├── sys_nonppc64x.go
│   ├── sys_openbsd_386.s
│   ├── sys_openbsd_amd64.s
│   ├── sys_openbsd_arm.s
│   ├── sys_openbsd_arm64.s
│   ├── sys_plan9_386.s
│   ├── sys_plan9_amd64.s
│   ├── sys_plan9_arm.s
│   ├── sys_ppc64x.go
│   ├── sys_s390x.go
│   ├── sys_solaris_amd64.s
│   ├── sys_wasm.go
│   ├── sys_wasm.s
│   ├── sys_windows_386.s
│   ├── sys_windows_amd64.s
│   ├── sys_windows_arm.s
│   ├── sys_x86.go
│   ├── syscall2_solaris.go
│   ├── syscall_aix.go
│   ├── syscall_nacl.h
│   ├── syscall_solaris.go
│   ├── syscall_windows.go
│   ├── syscall_windows_test.go
│   ├── testdata
│   │   ├── testprog
│   │   │   ├── abort.go
│   │   │   ├── badtraceback.go
│   │   │   ├── crash.go
│   │   │   ├── deadlock.go
│   │   │   ├── gc.go
│   │   │   ├── lockosthread.go
│   │   │   ├── main.go
│   │   │   ├── map.go
│   │   │   ├── memprof.go
│   │   │   ├── misc.go
│   │   │   ├── numcpu_freebsd.go
│   │   │   ├── panicrace.go
│   │   │   ├── signal.go
│   │   │   ├── sleep.go
│   │   │   ├── stringconcat.go
│   │   │   ├── syscall_windows.go
│   │   │   ├── syscalls.go
│   │   │   ├── syscalls_linux.go
│   │   │   ├── syscalls_none.go
│   │   │   ├── timeprof.go
│   │   │   └── traceback_ancestors.go
│   │   ├── testprogcgo
│   │   │   ├── aprof.go
│   │   │   ├── bigstack_windows.c
│   │   │   ├── bigstack_windows.go
│   │   │   ├── callback.go
│   │   │   ├── catchpanic.go
│   │   │   ├── cgo.go
│   │   │   ├── crash.go
│   │   │   ├── deadlock.go
│   │   │   ├── dll_windows.go
│   │   │   ├── dropm.go
│   │   │   ├── dropm_stub.go
│   │   │   ├── exec.go
│   │   │   ├── lockosthread.c
│   │   │   ├── lockosthread.go
│   │   │   ├── main.go
│   │   │   ├── numgoroutine.go
│   │   │   ├── pprof.go
│   │   │   ├── raceprof.go
│   │   │   ├── racesig.go
│   │   │   ├── sigpanic.go
│   │   │   ├── sigstack.go
│   │   │   ├── stack_windows.go
│   │   │   ├── threadpanic.go
│   │   │   ├── threadpanic_unix.c
│   │   │   ├── threadpanic_windows.c
│   │   │   ├── threadpprof.go
│   │   │   ├── threadprof.go
│   │   │   ├── traceback.go
│   │   │   ├── tracebackctxt.go
│   │   │   ├── tracebackctxt_c.c
│   │   │   └── windows
│   │   │       └── win.go
│   │   └── testprognet
│   │       ├── main.go
│   │       ├── net.go
│   │       ├── signal.go
│   │       └── signalexec.go
│   ├── textflag.h
│   ├── time.go
│   ├── timeasm.go
│   ├── timestub.go
│   ├── timestub2.go
│   ├── tls_arm.s
│   ├── tls_arm64.h
│   ├── tls_arm64.s
│   ├── tls_mips64x.s
│   ├── tls_mipsx.s
│   ├── tls_ppc64x.s
│   ├── tls_s390x.s
│   ├── trace
│   │   ├── annotation.go
│   │   ├── annotation_test.go
│   │   ├── example_test.go
│   │   ├── trace.go
│   │   ├── trace_stack_test.go
│   │   └── trace_test.go
│   ├── trace.go
│   ├── traceback.go
│   ├── treap_test.go
│   ├── type.go
│   ├── typekind.go
│   ├── utf8.go
│   ├── vdso_elf32.go
│   ├── vdso_elf64.go
│   ├── vdso_freebsd.go
│   ├── vdso_freebsd_arm.go
│   ├── vdso_freebsd_x86.go
│   ├── vdso_in_none.go
│   ├── vdso_linux.go
│   ├── vdso_linux_386.go
│   ├── vdso_linux_amd64.go
│   ├── vdso_linux_arm.go
│   ├── vdso_linux_arm64.go
│   ├── vdso_linux_ppc64x.go
│   ├── vlop_386.s
│   ├── vlop_arm.s
│   ├── vlop_arm_test.go
│   ├── vlrt.go
│   ├── wincallback.go
│   ├── write_err.go
│   ├── write_err_android.go
│   ├── zcallback_windows.go
│   ├── zcallback_windows.s
│   └── zcallback_windows_arm.s
├── sort
│   ├── example_interface_test.go
│   ├── example_keys_test.go
│   ├── example_multi_test.go
│   ├── example_search_test.go
│   ├── example_test.go
│   ├── example_wrapper_test.go
│   ├── export_test.go
│   ├── genzfunc.go
│   ├── search.go
│   ├── search_test.go
│   ├── slice.go
│   ├── slice_go113.go
│   ├── slice_go14.go
│   ├── slice_go18.go
│   ├── sort.go
│   ├── sort_test.go
│   └── zfuncversion.go
├── strconv
│   ├── atob.go
│   ├── atob_test.go
│   ├── atof.go
│   ├── atof_test.go
│   ├── atoi.go
│   ├── atoi_test.go
│   ├── decimal.go
│   ├── decimal_test.go
│   ├── doc.go
│   ├── example_test.go
│   ├── export_test.go
│   ├── extfloat.go
│   ├── fp_test.go
│   ├── ftoa.go
│   ├── ftoa_test.go
│   ├── internal_test.go
│   ├── isprint.go
│   ├── itoa.go
│   ├── itoa_test.go
│   ├── makeisprint.go
│   ├── quote.go
│   ├── quote_test.go
│   ├── strconv_test.go
│   └── testdata
│       └── testfp.txt
├── strings
│   ├── builder.go
│   ├── builder_test.go
│   ├── compare.go
│   ├── compare_test.go
│   ├── example_test.go
│   ├── export_test.go
│   ├── reader.go
│   ├── reader_test.go
│   ├── replace.go
│   ├── replace_test.go
│   ├── search.go
│   ├── search_test.go
│   ├── strings.go
│   └── strings_test.go
├── sync
│   ├── atomic
│   │   ├── asm.s
│   │   ├── atomic_test.go
│   │   ├── doc.go
│   │   ├── example_test.go
│   │   ├── race.s
│   │   ├── value.go
│   │   └── value_test.go
│   ├── cond.go
│   ├── cond_test.go
│   ├── example_pool_test.go
│   ├── example_test.go
│   ├── export_test.go
│   ├── map.go
│   ├── map_bench_test.go
│   ├── map_reference_test.go
│   ├── map_test.go
│   ├── mutex.go
│   ├── mutex_test.go
│   ├── once.go
│   ├── once_test.go
│   ├── pool.go
│   ├── pool_test.go
│   ├── poolqueue.go
│   ├── runtime.go
│   ├── runtime_sema_test.go
│   ├── rwmutex.go
│   ├── rwmutex_test.go
│   ├── waitgroup.go
│   └── waitgroup_test.go
├── syscall
│   ├── asm9_unix1_amd64.s
│   ├── asm9_unix2_amd64.s
│   ├── asm_aix_ppc64.s
│   ├── asm_darwin_386.s
│   ├── asm_darwin_amd64.s
│   ├── asm_darwin_arm.s
│   ├── asm_darwin_arm64.s
│   ├── asm_freebsd_arm.s
│   ├── asm_linux_386.s
│   ├── asm_linux_amd64.s
│   ├── asm_linux_arm.s
│   ├── asm_linux_arm64.s
│   ├── asm_linux_mips64x.s
│   ├── asm_linux_mipsx.s
│   ├── asm_linux_ppc64x.s
│   ├── asm_linux_s390x.s
│   ├── asm_nacl_386.s
│   ├── asm_nacl_amd64p32.s
│   ├── asm_nacl_arm.s
│   ├── asm_netbsd_arm.s
│   ├── asm_netbsd_arm64.s
│   ├── asm_openbsd_arm.s
│   ├── asm_openbsd_arm64.s
│   ├── asm_plan9_386.s
│   ├── asm_plan9_amd64.s
│   ├── asm_plan9_arm.s
│   ├── asm_solaris_amd64.s
│   ├── asm_unix_386.s
│   ├── asm_unix_amd64.s
│   ├── asm_windows.s
│   ├── bpf_bsd.go
│   ├── bpf_darwin.go
│   ├── const_plan9.go
│   ├── creds_test.go
│   ├── dir_plan9.go
│   ├── dirent.go
│   ├── dirent_bsd_test.go
│   ├── dll_windows.go
│   ├── endian_big.go
│   ├── endian_little.go
│   ├── env_plan9.go
│   ├── env_unix.go
│   ├── env_windows.go
│   ├── errors_plan9.go
│   ├── exec_aix_test.go
│   ├── exec_bsd.go
│   ├── exec_darwin.go
│   ├── exec_libc.go
│   ├── exec_linux.go
│   ├── exec_linux_test.go
│   ├── exec_plan9.go
│   ├── exec_solaris_test.go
│   ├── exec_unix.go
│   ├── exec_unix_test.go
│   ├── exec_windows.go
│   ├── export_freebsd_test.go
│   ├── export_linux_test.go
│   ├── export_test.go
│   ├── export_unix_test.go
│   ├── fd_nacl.go
│   ├── flock.go
│   ├── flock_aix.go
│   ├── flock_darwin.go
│   ├── flock_linux_32bit.go
│   ├── forkpipe.go
│   ├── forkpipe2.go
│   ├── fs_js.go
│   ├── fs_nacl.go
│   ├── getdirentries_test.go
│   ├── js
│   │   ├── func.go
│   │   ├── js.go
│   │   ├── js_js.s
│   │   └── js_test.go
│   ├── lsf_linux.go
│   ├── mkall.sh
│   ├── mkasm_darwin.go
│   ├── mkerrors.sh
│   ├── mkpost.go
│   ├── mksyscall.pl
│   ├── mksyscall_libc.pl
│   ├── mksyscall_windows.go
│   ├── mksysctl_openbsd.pl
│   ├── mksysnum_darwin.pl
│   ├── mksysnum_dragonfly.pl
│   ├── mksysnum_freebsd.pl
│   ├── mksysnum_linux.pl
│   ├── mksysnum_netbsd.pl
│   ├── mksysnum_openbsd.pl
│   ├── mksysnum_plan9.sh
│   ├── mmap_unix_test.go
│   ├── msan.go
│   ├── msan0.go
│   ├── net.go
│   ├── net_js.go
│   ├── net_nacl.go
│   ├── netlink_linux.go
│   ├── pwd_plan9.go
│   ├── route_bsd.go
│   ├── route_darwin.go
│   ├── route_dragonfly.go
│   ├── route_freebsd.go
│   ├── route_freebsd_32bit.go
│   ├── route_freebsd_64bit.go
│   ├── route_netbsd.go
│   ├── route_openbsd.go
│   ├── security_windows.go
│   ├── setuidgid_32_linux.go
│   ├── setuidgid_linux.go
│   ├── sockcmsg_linux.go
│   ├── sockcmsg_unix.go
│   ├── str.go
│   ├── syscall.go
│   ├── syscall_aix.go
│   ├── syscall_aix_ppc64.go
│   ├── syscall_bsd.go
│   ├── syscall_bsd_test.go
│   ├── syscall_darwin.go
│   ├── syscall_darwin_386.go
│   ├── syscall_darwin_amd64.go
│   ├── syscall_darwin_arm.go
│   ├── syscall_darwin_arm64.go
│   ├── syscall_dragonfly.go
│   ├── syscall_dragonfly_amd64.go
│   ├── syscall_freebsd.go
│   ├── syscall_freebsd_386.go
│   ├── syscall_freebsd_amd64.go
│   ├── syscall_freebsd_arm.go
│   ├── syscall_freebsd_test.go
│   ├── syscall_getwd_bsd.go
│   ├── syscall_js.go
│   ├── syscall_linux.go
│   ├── syscall_linux_386.go
│   ├── syscall_linux_amd64.go
│   ├── syscall_linux_arm.go
│   ├── syscall_linux_arm64.go
│   ├── syscall_linux_mips64x.go
│   ├── syscall_linux_mipsx.go
│   ├── syscall_linux_ppc64x.go
│   ├── syscall_linux_s390x.go
│   ├── syscall_linux_test.go
│   ├── syscall_nacl.go
│   ├── syscall_nacl_386.go
│   ├── syscall_nacl_amd64p32.go
│   ├── syscall_nacl_arm.go
│   ├── syscall_netbsd.go
│   ├── syscall_netbsd_386.go
│   ├── syscall_netbsd_amd64.go
│   ├── syscall_netbsd_arm.go
│   ├── syscall_netbsd_arm64.go
│   ├── syscall_openbsd.go
│   ├── syscall_openbsd_386.go
│   ├── syscall_openbsd_amd64.go
│   ├── syscall_openbsd_arm.go
│   ├── syscall_openbsd_arm64.go
│   ├── syscall_plan9.go
│   ├── syscall_plan9_test.go
│   ├── syscall_ptrace_test.go
│   ├── syscall_solaris.go
│   ├── syscall_solaris_amd64.go
│   ├── syscall_test.go
│   ├── syscall_unix.go
│   ├── syscall_unix_test.go
│   ├── syscall_windows.go
│   ├── syscall_windows_386.go
│   ├── syscall_windows_amd64.go
│   ├── syscall_windows_test.go
│   ├── tables_nacljs.go
│   ├── time_nacl_386.s
│   ├── time_nacl_amd64p32.s
│   ├── time_nacl_arm.s
│   ├── timestruct.go
│   ├── types_aix.go
│   ├── types_darwin.go
│   ├── types_dragonfly.go
│   ├── types_freebsd.go
│   ├── types_linux.go
│   ├── types_netbsd.go
│   ├── types_openbsd.go
│   ├── types_solaris.go
│   ├── types_windows.go
│   ├── types_windows_386.go
│   ├── types_windows_amd64.go
│   ├── types_windows_arm.go
│   ├── unzip_nacl.go
│   ├── zerrors_aix_ppc64.go
│   ├── zerrors_darwin_386.go
│   ├── zerrors_darwin_amd64.go
│   ├── zerrors_darwin_arm.go
│   ├── zerrors_darwin_arm64.go
│   ├── zerrors_dragonfly_amd64.go
│   ├── zerrors_freebsd_386.go
│   ├── zerrors_freebsd_amd64.go
│   ├── zerrors_freebsd_arm.go
│   ├── zerrors_linux_386.go
│   ├── zerrors_linux_amd64.go
│   ├── zerrors_linux_arm.go
│   ├── zerrors_linux_arm64.go
│   ├── zerrors_linux_mips.go
│   ├── zerrors_linux_mips64.go
│   ├── zerrors_linux_mips64le.go
│   ├── zerrors_linux_mipsle.go
│   ├── zerrors_linux_ppc64.go
│   ├── zerrors_linux_ppc64le.go
│   ├── zerrors_linux_s390x.go
│   ├── zerrors_netbsd_386.go
│   ├── zerrors_netbsd_amd64.go
│   ├── zerrors_netbsd_arm.go
│   ├── zerrors_netbsd_arm64.go
│   ├── zerrors_openbsd_386.go
│   ├── zerrors_openbsd_amd64.go
│   ├── zerrors_openbsd_arm.go
│   ├── zerrors_openbsd_arm64.go
│   ├── zerrors_solaris_amd64.go
│   ├── zerrors_windows.go
│   ├── zerrors_windows_386.go
│   ├── zerrors_windows_amd64.go
│   ├── zsyscall_aix_ppc64.go
│   ├── zsyscall_darwin_386.go
│   ├── zsyscall_darwin_386.s
│   ├── zsyscall_darwin_amd64.go
│   ├── zsyscall_darwin_amd64.s
│   ├── zsyscall_darwin_arm.go
│   ├── zsyscall_darwin_arm.s
│   ├── zsyscall_darwin_arm64.go
│   ├── zsyscall_darwin_arm64.s
│   ├── zsyscall_dragonfly_amd64.go
│   ├── zsyscall_freebsd_386.go
│   ├── zsyscall_freebsd_amd64.go
│   ├── zsyscall_freebsd_arm.go
│   ├── zsyscall_linux_386.go
│   ├── zsyscall_linux_amd64.go
│   ├── zsyscall_linux_arm.go
│   ├── zsyscall_linux_arm64.go
│   ├── zsyscall_linux_mips.go
│   ├── zsyscall_linux_mips64.go
│   ├── zsyscall_linux_mips64le.go
│   ├── zsyscall_linux_mipsle.go
│   ├── zsyscall_linux_ppc64.go
│   ├── zsyscall_linux_ppc64le.go
│   ├── zsyscall_linux_s390x.go
│   ├── zsyscall_nacl_386.go
│   ├── zsyscall_nacl_amd64p32.go
│   ├── zsyscall_nacl_arm.go
│   ├── zsyscall_netbsd_386.go
│   ├── zsyscall_netbsd_amd64.go
│   ├── zsyscall_netbsd_arm.go
│   ├── zsyscall_netbsd_arm64.go
│   ├── zsyscall_openbsd_386.go
│   ├── zsyscall_openbsd_amd64.go
│   ├── zsyscall_openbsd_arm.go
│   ├── zsyscall_openbsd_arm64.go
│   ├── zsyscall_plan9_386.go
│   ├── zsyscall_plan9_amd64.go
│   ├── zsyscall_plan9_arm.go
│   ├── zsyscall_solaris_amd64.go
│   ├── zsyscall_windows.go
│   ├── zsysctl_openbsd.go
│   ├── zsysnum_darwin_386.go
│   ├── zsysnum_darwin_amd64.go
│   ├── zsysnum_darwin_arm.go
│   ├── zsysnum_darwin_arm64.go
│   ├── zsysnum_dragonfly_amd64.go
│   ├── zsysnum_freebsd_386.go
│   ├── zsysnum_freebsd_amd64.go
│   ├── zsysnum_freebsd_arm.go
│   ├── zsysnum_linux_386.go
│   ├── zsysnum_linux_amd64.go
│   ├── zsysnum_linux_arm.go
│   ├── zsysnum_linux_arm64.go
│   ├── zsysnum_linux_mips.go
│   ├── zsysnum_linux_mips64.go
│   ├── zsysnum_linux_mips64le.go
│   ├── zsysnum_linux_mipsle.go
│   ├── zsysnum_linux_ppc64.go
│   ├── zsysnum_linux_ppc64le.go
│   ├── zsysnum_linux_s390x.go
│   ├── zsysnum_netbsd_386.go
│   ├── zsysnum_netbsd_amd64.go
│   ├── zsysnum_netbsd_arm.go
│   ├── zsysnum_netbsd_arm64.go
│   ├── zsysnum_openbsd_386.go
│   ├── zsysnum_openbsd_amd64.go
│   ├── zsysnum_openbsd_arm.go
│   ├── zsysnum_openbsd_arm64.go
│   ├── zsysnum_plan9.go
│   ├── zsysnum_solaris_amd64.go
│   ├── zsysnum_windows_386.go
│   ├── zsysnum_windows_amd64.go
│   ├── ztypes_aix_ppc64.go
│   ├── ztypes_darwin_386.go
│   ├── ztypes_darwin_amd64.go
│   ├── ztypes_darwin_arm.go
│   ├── ztypes_darwin_arm64.go
│   ├── ztypes_dragonfly_amd64.go
│   ├── ztypes_freebsd_386.go
│   ├── ztypes_freebsd_amd64.go
│   ├── ztypes_freebsd_arm.go
│   ├── ztypes_linux_386.go
│   ├── ztypes_linux_amd64.go
│   ├── ztypes_linux_arm.go
│   ├── ztypes_linux_arm64.go
│   ├── ztypes_linux_mips.go
│   ├── ztypes_linux_mips64.go
│   ├── ztypes_linux_mips64le.go
│   ├── ztypes_linux_mipsle.go
│   ├── ztypes_linux_ppc64.go
│   ├── ztypes_linux_ppc64le.go
│   ├── ztypes_linux_s390x.go
│   ├── ztypes_netbsd_386.go
│   ├── ztypes_netbsd_amd64.go
│   ├── ztypes_netbsd_arm.go
│   ├── ztypes_netbsd_arm64.go
│   ├── ztypes_openbsd_386.go
│   ├── ztypes_openbsd_amd64.go
│   ├── ztypes_openbsd_arm.go
│   ├── ztypes_openbsd_arm64.go
│   └── ztypes_solaris_amd64.go
├── testdata
│   └── Isaac.Newton-Opticks.txt
├── testing
│   ├── allocs.go
│   ├── allocs_test.go
│   ├── benchmark.go
│   ├── benchmark_test.go
│   ├── cover.go
│   ├── example.go
│   ├── export_test.go
│   ├── helper_test.go
│   ├── helperfuncs_test.go
│   ├── internal
│   │   └── testdeps
│   │       └── deps.go
│   ├── iotest
│   │   ├── logger.go
│   │   ├── reader.go
│   │   └── writer.go
│   ├── match.go
│   ├── match_test.go
│   ├── quick
│   │   ├── quick.go
│   │   └── quick_test.go
│   ├── run_example.go
│   ├── run_example_js.go
│   ├── sub_test.go
│   ├── testing.go
│   └── testing_test.go
├── text
│   ├── scanner
│   │   ├── example_test.go
│   │   ├── scanner.go
│   │   └── scanner_test.go
│   ├── tabwriter
│   │   ├── example_test.go
│   │   ├── tabwriter.go
│   │   └── tabwriter_test.go
│   └── template
│       ├── doc.go
│       ├── example_test.go
│       ├── examplefiles_test.go
│       ├── examplefunc_test.go
│       ├── exec.go
│       ├── exec_test.go
│       ├── funcs.go
│       ├── helper.go
│       ├── multi_test.go
│       ├── option.go
│       ├── parse
│       │   ├── lex.go
│       │   ├── lex_test.go
│       │   ├── node.go
│       │   ├── parse.go
│       │   └── parse_test.go
│       ├── template.go
│       └── testdata
│           ├── file1.tmpl
│           ├── file2.tmpl
│           ├── tmpl1.tmpl
│           └── tmpl2.tmpl
├── time
│   ├── example_test.go
│   ├── export_android_test.go
│   ├── export_test.go
│   ├── export_windows_test.go
│   ├── format.go
│   ├── format_test.go
│   ├── genzabbrs.go
│   ├── internal_test.go
│   ├── mono_test.go
│   ├── sleep.go
│   ├── sleep_test.go
│   ├── sys_plan9.go
│   ├── sys_unix.go
│   ├── sys_windows.go
│   ├── tick.go
│   ├── tick_test.go
│   ├── time.go
│   ├── time_test.go
│   ├── zoneinfo.go
│   ├── zoneinfo_abbrs_windows.go
│   ├── zoneinfo_android.go
│   ├── zoneinfo_android_test.go
│   ├── zoneinfo_ios.go
│   ├── zoneinfo_js.go
│   ├── zoneinfo_plan9.go
│   ├── zoneinfo_read.go
│   ├── zoneinfo_test.go
│   ├── zoneinfo_unix.go
│   ├── zoneinfo_windows.go
│   └── zoneinfo_windows_test.go
├── unicode
│   ├── casetables.go
│   ├── digit.go
│   ├── digit_test.go
│   ├── example_test.go
│   ├── graphic.go
│   ├── graphic_test.go
│   ├── letter.go
│   ├── letter_test.go
│   ├── script_test.go
│   ├── tables.go
│   ├── utf16
│   │   ├── export_test.go
│   │   ├── utf16.go
│   │   └── utf16_test.go
│   └── utf8
│       ├── example_test.go
│       ├── utf8.go
│       └── utf8_test.go
├── unsafe
│   └── unsafe.go
└── vendor
    ├── golang.org
    │   └── x
    │       ├── crypto
    │       │   ├── AUTHORS
    │       │   ├── CONTRIBUTORS
    │       │   ├── LICENSE
    │       │   ├── PATENTS
    │       │   ├── chacha20poly1305
    │       │   │   ├── chacha20poly1305.go
    │       │   │   ├── chacha20poly1305_amd64.go
    │       │   │   ├── chacha20poly1305_amd64.s
    │       │   │   ├── chacha20poly1305_generic.go
    │       │   │   ├── chacha20poly1305_noasm.go
    │       │   │   └── xchacha20poly1305.go
    │       │   ├── cryptobyte
    │       │   │   ├── asn1
    │       │   │   │   └── asn1.go
    │       │   │   ├── asn1.go
    │       │   │   ├── builder.go
    │       │   │   └── string.go
    │       │   ├── curve25519
    │       │   │   ├── const_amd64.h
    │       │   │   ├── const_amd64.s
    │       │   │   ├── cswap_amd64.s
    │       │   │   ├── curve25519.go
    │       │   │   ├── doc.go
    │       │   │   ├── freeze_amd64.s
    │       │   │   ├── ladderstep_amd64.s
    │       │   │   ├── mont25519_amd64.go
    │       │   │   ├── mul_amd64.s
    │       │   │   └── square_amd64.s
    │       │   ├── hkdf
    │       │   │   └── hkdf.go
    │       │   ├── internal
    │       │   │   ├── chacha20
    │       │   │   │   ├── asm_arm64.s
    │       │   │   │   ├── asm_ppc64le.s
    │       │   │   │   ├── chacha_arm64.go
    │       │   │   │   ├── chacha_generic.go
    │       │   │   │   ├── chacha_noasm.go
    │       │   │   │   ├── chacha_ppc64le.go
    │       │   │   │   ├── chacha_s390x.go
    │       │   │   │   ├── chacha_s390x.s
    │       │   │   │   └── xor.go
    │       │   │   └── subtle
    │       │   │       ├── aliasing.go
    │       │   │       └── aliasing_appengine.go
    │       │   └── poly1305
    │       │       ├── mac_noasm.go
    │       │       ├── poly1305.go
    │       │       ├── sum_amd64.go
    │       │       ├── sum_amd64.s
    │       │       ├── sum_arm.go
    │       │       ├── sum_arm.s
    │       │       ├── sum_generic.go
    │       │       ├── sum_noasm.go
    │       │       ├── sum_ppc64le.go
    │       │       ├── sum_ppc64le.s
    │       │       ├── sum_s390x.go
    │       │       ├── sum_s390x.s
    │       │       └── sum_vmsl_s390x.s
    │       ├── net
    │       │   ├── AUTHORS
    │       │   ├── CONTRIBUTORS
    │       │   ├── LICENSE
    │       │   ├── PATENTS
    │       │   ├── dns
    │       │   │   └── dnsmessage
    │       │   │       └── message.go
    │       │   ├── http
    │       │   │   ├── httpguts
    │       │   │   │   ├── guts.go
    │       │   │   │   └── httplex.go
    │       │   │   └── httpproxy
    │       │   │       └── proxy.go
    │       │   ├── http2
    │       │   │   └── hpack
    │       │   │       ├── encode.go
    │       │   │       ├── hpack.go
    │       │   │       ├── huffman.go
    │       │   │       └── tables.go
    │       │   ├── idna
    │       │   │   ├── idna10.0.0.go
    │       │   │   ├── idna9.0.0.go
    │       │   │   ├── punycode.go
    │       │   │   ├── tables10.0.0.go
    │       │   │   ├── tables11.0.0.go
    │       │   │   ├── tables9.0.0.go
    │       │   │   ├── trie.go
    │       │   │   └── trieval.go
    │       │   ├── lif
    │       │   │   ├── address.go
    │       │   │   ├── binary.go
    │       │   │   ├── lif.go
    │       │   │   ├── link.go
    │       │   │   ├── sys.go
    │       │   │   ├── sys_solaris_amd64.s
    │       │   │   ├── syscall.go
    │       │   │   └── zsys_solaris_amd64.go
    │       │   ├── nettest
    │       │   │   ├── conntest.go
    │       │   │   ├── nettest.go
    │       │   │   ├── nettest_stub.go
    │       │   │   ├── nettest_unix.go
    │       │   │   └── nettest_windows.go
    │       │   └── route
    │       │       ├── address.go
    │       │       ├── binary.go
    │       │       ├── empty.s
    │       │       ├── interface.go
    │       │       ├── interface_announce.go
    │       │       ├── interface_classic.go
    │       │       ├── interface_freebsd.go
    │       │       ├── interface_multicast.go
    │       │       ├── interface_openbsd.go
    │       │       ├── message.go
    │       │       ├── route.go
    │       │       ├── route_classic.go
    │       │       ├── route_openbsd.go
    │       │       ├── sys.go
    │       │       ├── sys_darwin.go
    │       │       ├── sys_dragonfly.go
    │       │       ├── sys_freebsd.go
    │       │       ├── sys_netbsd.go
    │       │       ├── sys_openbsd.go
    │       │       ├── syscall.go
    │       │       ├── syscall_go1_11_darwin.go
    │       │       ├── syscall_go1_12_darwin.go
    │       │       ├── zsys_darwin.go
    │       │       ├── zsys_dragonfly.go
    │       │       ├── zsys_freebsd_386.go
    │       │       ├── zsys_freebsd_amd64.go
    │       │       ├── zsys_freebsd_arm.go
    │       │       ├── zsys_freebsd_arm64.go
    │       │       ├── zsys_netbsd.go
    │       │       └── zsys_openbsd.go
    │       ├── sys
    │       │   ├── AUTHORS
    │       │   ├── CONTRIBUTORS
    │       │   ├── LICENSE
    │       │   ├── PATENTS
    │       │   └── cpu
    │       │       ├── asm_aix_ppc64.s
    │       │       ├── byteorder.go
    │       │       ├── cpu.go
    │       │       ├── cpu_aix_ppc64.go
    │       │       ├── cpu_arm.go
    │       │       ├── cpu_gc_s390x.go
    │       │       ├── cpu_gc_x86.go
    │       │       ├── cpu_gccgo.c
    │       │       ├── cpu_gccgo.go
    │       │       ├── cpu_gccgo_s390x.go
    │       │       ├── cpu_linux.go
    │       │       ├── cpu_linux_arm64.go
    │       │       ├── cpu_linux_ppc64x.go
    │       │       ├── cpu_linux_s390x.go
    │       │       ├── cpu_mips64x.go
    │       │       ├── cpu_mipsx.go
    │       │       ├── cpu_other_arm64.go
    │       │       ├── cpu_s390x.s
    │       │       ├── cpu_wasm.go
    │       │       ├── cpu_x86.go
    │       │       ├── cpu_x86.s
    │       │       └── syscall_aix_ppc64_gc.go
    │       └── text
    │           ├── AUTHORS
    │           ├── CONTRIBUTORS
    │           ├── LICENSE
    │           ├── PATENTS
    │           ├── secure
    │           │   └── bidirule
    │           │       ├── bidirule.go
    │           │       ├── bidirule10.0.0.go
    │           │       └── bidirule9.0.0.go
    │           ├── transform
    │           │   └── transform.go
    │           └── unicode
    │               ├── bidi
    │               │   ├── bidi.go
    │               │   ├── bracket.go
    │               │   ├── core.go
    │               │   ├── prop.go
    │               │   ├── tables10.0.0.go
    │               │   ├── tables11.0.0.go
    │               │   ├── tables9.0.0.go
    │               │   └── trieval.go
    │               └── norm
    │                   ├── composition.go
    │                   ├── forminfo.go
    │                   ├── input.go
    │                   ├── iter.go
    │                   ├── normalize.go
    │                   ├── readwriter.go
    │                   ├── tables10.0.0.go
    │                   ├── tables11.0.0.go
    │                   ├── tables9.0.0.go
    │                   ├── transform.go
    │                   └── trie.go
    └── modules.txt

771 directories, 6012 files
