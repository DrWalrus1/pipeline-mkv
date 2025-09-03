interface MakeMkvService {
  GetDiscInfo: () => Promise<string>
}

class MockMakeMkvService implements MakeMkvService {
  public GetDiscInfo(): Promise<string> {
    return new Promise<string>((resolve, reject) => {
      resolve("Hallo")
    })
  }
}
