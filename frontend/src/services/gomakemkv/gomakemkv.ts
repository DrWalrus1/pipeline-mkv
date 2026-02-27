import type { DiscInfo } from "@/domain/disc_info"

export interface MakeMkvService {
  GetInsertedDiscInfo: () => Promise<string>
  GetRecentDiscInfos: (limit: number, offset: number) => Promise<DiscInfo[]>
}

export type MakeMkvServiceProvider = MakeMkvService | undefined

export function GetMakeMkvService(x: MakeMkvServiceProvider): asserts x is MakeMkvService {
  if (x == null) {
    throw new Error("MakeMkvService has not been provided by parent higher in component tree")
  }
}

