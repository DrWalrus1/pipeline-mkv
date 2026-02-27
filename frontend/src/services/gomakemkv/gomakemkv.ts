export interface MakeMkvService {
  GetDiscInfo: () => Promise<string>
}

export type MakeMkvServiceProvider = MakeMkvService | undefined

export function GetMakeMkvService(x: MakeMkvServiceProvider): asserts x is MakeMkvService {
  if (x == null) {
    throw new Error("MakeMkvService has not been provided by parent higher in component tree")
  }
}

class MockMakeMkvService implements MakeMkvService {
  public GetDiscInfo(): Promise<string> {
    return new Promise<string>((resolve, reject) => {
      resolve("Hallo")
    })
  }
}
