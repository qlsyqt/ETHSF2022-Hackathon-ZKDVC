

可信设置代码
```

snarkjs powersoftau new bn128 16 pot_0000.ptau -v
snarkjs powersoftau contribute pot_0000.ptau pot_0001.ptau --name="First contribution" -v
snarkjs powersoftau prepare phase2 pot_0001.ptau pot_final.ptau -v
snarkjs groth16 setup stateTransition.r1cs pot_final.ptau circuit_0000.zkey
snarkjs zkey contribute circuit_0000.zkey stateTransition.zkey --name="First contribution" -v
snarkjs zkey export verificationkey stateTransition.zkey verification_key.json

snarkjs groth16 prove stateTransition.zkey witness.wtns proof.json public.json

snarkjs groth16 verify verification_key.json public.json proof.json

snarkjs zkey export solidityverifier circuit_0001.zkey verifier.sol
```